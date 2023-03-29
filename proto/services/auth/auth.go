package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/galexrt/arpanet/pkg/auth"
	"github.com/galexrt/arpanet/pkg/perms"
	"github.com/galexrt/arpanet/proto/resources/jobs"
	users "github.com/galexrt/arpanet/proto/resources/users"
	"github.com/galexrt/arpanet/query/arpanet/model"
	"github.com/galexrt/arpanet/query/arpanet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/slices"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	account   = table.ArpanetAccounts
	user      = table.Users.AS("user")
	js        = table.Jobs
	jobGrades = table.JobGrades
	jobProps  = table.ArpanetJobProps
)

var (
	AccountCreateFailedErr = status.Error(codes.InvalidArgument, "Please check the token you have typed in!")
	InvalidLoginErr        = status.Error(codes.InvalidArgument, "Wrong username or password.")
	NoPasswordErr          = status.Error(codes.InvalidArgument, "You need to use your registration token to create an account first.")
	NoCharacterFoundErr    = status.Error(codes.NotFound, "No character(s) found for your account.")
	GenericLoginErr        = status.Error(codes.Internal, "Failed to login you in, please try again.")
	UnableToChooseCharErr  = status.Error(codes.PermissionDenied, "You don't have permission to select this character!")
)

type Server struct {
	AuthServiceServer

	db   *sql.DB
	auth *auth.GRPCAuth
	tm   *auth.TokenManager
	p    perms.Permissions
}

func NewServer(db *sql.DB, auth *auth.GRPCAuth, tm *auth.TokenManager, p perms.Permissions) *Server {
	return &Server{
		db:   db,
		auth: auth,
		tm:   tm,
		p:    p,
	}
}

// AuthFuncOverride is called instead of exampleAuthFunc
func (s *Server) AuthFuncOverride(ctx context.Context, fullMethod string) (context.Context, error) {
	// Skip authentication for the anon accessible endpoints
	if fullMethod == "/services.auth.AuthService/CreateAccount" ||
		fullMethod == "/services.auth.AuthService/Login" ||
		fullMethod == "/services.auth.AuthService/Logout" {
		return ctx, nil
	}

	return s.auth.GRPCAuthFunc(ctx, fullMethod)
}

func (s *Server) PermissionUnaryFuncOverride(ctx context.Context, info *grpc.UnaryServerInfo) (context.Context, error) {
	// Skip permission check for the auth services
	return ctx, nil
}

func (s *Server) createTokenFromAccountAndChar(account *model.ArpanetAccounts, activeChar *users.User) (string, error) {
	claims := &auth.CitizenInfoClaims{
		AccountID: account.ID,
		Username:  *account.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "arpanet",
			Subject:   account.License,
			ID:        strconv.FormatUint(uint64(account.ID), 10),
			Audience:  []string{"arpanet"},
		},
	}

	if activeChar != nil {
		claims.ActiveCharID = activeChar.UserId
		claims.ActiveCharJob = activeChar.Job
		claims.ActiveCharJobGrade = activeChar.JobGrade
	} else {
		claims.ActiveCharID = 0
		claims.ActiveCharJob = ""
		claims.ActiveCharJobGrade = 0
	}

	return s.tm.NewWithClaims(claims)
}

func (s *Server) CreateAccount(ctx context.Context, req *CreateAccountRequest) (*CreateAccountResponse, error) {
	acc, err := s.getAccountFromDB(ctx, jet.AND(
		account.Password.IS_NULL(),
		account.RegToken.EQ(jet.String(req.RegToken)),
	))
	if err != nil {
		return nil, AccountCreateFailedErr
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, AccountCreateFailedErr
	}

	stmt := account.
		UPDATE().
		SET(
			account.Username.SET(jet.String(req.Username)),
			account.Password.SET(jet.String(string(hashedPassword))),
		).
		WHERE(
			jet.AND(
				account.ID.EQ(jet.Uint64(acc.ID)),
				account.Password.IS_NULL(),
				account.RegToken.EQ(jet.String(req.RegToken)),
			),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, AccountCreateFailedErr
	}

	return &CreateAccountResponse{}, nil
}

func (s *Server) getAccountFromDB(ctx context.Context, condition jet.BoolExpression) (*model.ArpanetAccounts, error) {
	stmt := account.
		SELECT(
			account.AllColumns,
		).
		FROM(account).
		WHERE(
			account.Enabled.IS_TRUE().
				AND(condition),
		).
		LIMIT(1)

	var acc model.ArpanetAccounts
	if err := stmt.QueryContext(ctx, s.db, &acc); err != nil {
		return nil, err
	}

	return &acc, nil
}

func (s *Server) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	account, err := s.getAccountFromDB(ctx, account.Username.EQ(jet.String(req.Username)))
	if err != nil {
		if errors.Is(qrm.ErrNoRows, err) {
			return nil, InvalidLoginErr
		}

		return nil, err
	}

	// No password set
	if account.Password == nil {
		return nil, NoPasswordErr
	}

	// Password check logic
	if err := bcrypt.CompareHashAndPassword([]byte(*account.Password), []byte(req.Password)); err != nil {
		return nil, InvalidLoginErr
	}

	token, err := s.createTokenFromAccountAndChar(account, nil)
	if err != nil {
		return nil, InvalidLoginErr
	}

	return &LoginResponse{
		Token: token,
	}, nil
}

func (s *Server) GetCharacters(ctx context.Context, req *GetCharactersRequest) (*GetCharactersResponse, error) {
	claims, err := s.tm.ParseWithClaims(auth.MustGetTokenFromGRPCContext(ctx))
	if err != nil {
		return nil, err
	}

	// Load chars from database
	stmt := user.
		SELECT(
			user.ID,
			user.Identifier,
			user.Job,
			js.Label.AS("user.job_label"),
			user.JobGrade,
			jobGrades.Label.AS("user.job_grade_label"),
			user.Firstname,
			user.Lastname,
			user.Dateofbirth,
			user.Sex,
			user.Height,
			user.PhoneNumber,
			user.Visum,
			user.Playtime,
		).
		FROM(user.
			LEFT_JOIN(js,
				js.Name.EQ(user.Job),
			).
			LEFT_JOIN(jobGrades,
				jet.AND(
					jobGrades.Grade.EQ(user.JobGrade),
					jobGrades.JobName.EQ(user.Job),
				),
			),
		).
		WHERE(
			user.Identifier.LIKE(jet.String(buildCharSearchIdentifier(claims.Subject))),
		).
		ORDER_BY(user.ID).
		LIMIT(10)

	resp := &GetCharactersResponse{}
	if err := stmt.QueryContext(ctx, s.db, &resp.Chars); err != nil {
		if errors.Is(qrm.ErrNoRows, err) {
			return nil, NoCharacterFoundErr
		}
		return nil, err
	}

	return resp, nil
}

func buildCharSearchIdentifier(license string) string {
	return "char%:" + license
}

func (s *Server) getCharacter(ctx context.Context, charId int32) (*users.User, *jobs.JobProps, error) {
	stmt := user.
		SELECT(
			user.ID,
			user.Identifier,
			user.Job,
			user.JobGrade,
			js.Label.AS("user.job_label"),
			jobGrades.Label.AS("user.job_grade_label"),
			jobProps.Theme,
		).
		FROM(
			user.
				LEFT_JOIN(js,
					js.Name.EQ(user.Job),
				).
				LEFT_JOIN(jobGrades,
					jet.AND(
						jobGrades.Grade.EQ(user.JobGrade),
						jobGrades.JobName.EQ(user.Job),
					),
				).
				LEFT_JOIN(jobProps,
					jobProps.Job.EQ(js.Name),
				),
		).
		WHERE(
			user.ID.EQ(jet.Int32(charId)),
		).
		LIMIT(1)

	var dest struct {
		users.User
		JobProps jobs.JobProps
	}
	if err := stmt.QueryContext(ctx, s.db, &dest); err != nil {
		if errors.Is(qrm.ErrNoRows, err) {
			return nil, nil, NoCharacterFoundErr
		}
		return nil, nil, err
	}

	return &dest.User, &dest.JobProps, nil
}

func (s *Server) ChooseCharacter(ctx context.Context, req *ChooseCharacterRequest) (*ChooseCharacterResponse, error) {
	claims, err := s.tm.ParseWithClaims(auth.MustGetTokenFromGRPCContext(ctx))
	if err != nil {
		return nil, err
	}

	char, jProps, err := s.getCharacter(ctx, req.CharId)
	if err != nil {
		return nil, err
	}

	// Make sure the user isn't sending us a different char ID than their own
	if !strings.Contains(char.Identifier, claims.Subject) {
		return nil, UnableToChooseCharErr
	}

	// Load account data for token creation
	account, err := s.getAccountFromDB(ctx, account.Username.EQ(jet.String(claims.Username)))
	if err != nil {
		return nil, err
	}

	token, err := s.createTokenFromAccountAndChar(account, char)
	if err != nil {
		return nil, err
	}

	if err := s.ensureUserHasRole(char.UserId, char.Job, char.JobGrade); err != nil {
		return nil, err
	}

	// Load permissions of user
	perms, err := s.p.GetAllPermissionsOfUser(char.UserId)
	if err != nil {
		return nil, err
	}

	if len(perms) == 0 {
		return nil, UnableToChooseCharErr
	}

	return &ChooseCharacterResponse{
		Token:       token,
		Permissions: perms.GuardNames(),
		JobProps:    jProps,
	}, nil
}

// Make sure the user is only in their current job role
func (s *Server) ensureUserHasRole(userId int32, job string, jobGrade int32) error {
	jobRoleKey := fmt.Sprintf("job-%s-", job)
	jobRolesModels, err := s.p.GetRoles(jobRoleKey)
	if err != nil {
		return err
	}

	if len(jobRolesModels) == 0 {
		return nil
	}

	jobRoles := jobRolesModels.GuardNames()

	var jobRole string
	// Try to see if there is an exact role match for the job and grade of the user
	index := slices.Index(jobRoles, perms.GetRoleName(job, jobGrade))
	if index < 0 {
		for i := len(jobRoles) - 1; i >= 0; i-- {
			_, gradeString, _ := strings.Cut(jobRoles[i], jobRoleKey)
			grade, err := strconv.Atoi(gradeString)
			if err != nil {
				continue
			}

			// Never upgrade an user's role
			if grade > int(jobGrade) {
				continue
			}

			index = i

			if grade <= int(jobGrade) {
				break
			}
		}
	}

	if index < 0 {
		return nil
	}

	jobRole = jobRoles[index]

	ps, err := s.p.GetUserRoles(userId)
	if err != nil {
		return err
	}

	rolesToRemove := []string{}
	for _, name := range ps.GuardNames() {
		// Ignore roles that don't start with the `job-` prefix
		if !strings.HasPrefix(name, "job-") {
			continue
		}

		if name != jobRole {
			rolesToRemove = append(rolesToRemove, name)
		}
	}

	if err := s.p.RemoveUserRoles(userId, rolesToRemove...); err != nil {
		return err
	}

	if err := s.p.AddUserRoles(userId, jobRole); err != nil {
		return err
	}

	return nil
}

func (s *Server) Logout(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	return &LogoutResponse{
		Success: true,
	}, nil
}

func (s *Server) SetJob(ctx context.Context, req *SetJobRequest) (*SetJobResponse, error) {
	claims, err := s.tm.ParseWithClaims(auth.MustGetTokenFromGRPCContext(ctx))
	if err != nil {
		return nil, err
	}

	char, jProps, err := s.getCharacter(ctx, claims.ActiveCharID)
	if err != nil {
		return nil, err
	}

	char.Job = req.Job
	char.JobGrade = req.JobGrade

	// Load account data for token creation
	account, err := s.getAccountFromDB(ctx, account.Username.EQ(jet.String(claims.Username)))
	if err != nil {
		return nil, err
	}

	token, err := s.createTokenFromAccountAndChar(account, char)
	if err != nil {
		return nil, err
	}

	return &SetJobResponse{
		Token:    token,
		JobProps: jProps,
		Char:     char,
	}, nil
}
