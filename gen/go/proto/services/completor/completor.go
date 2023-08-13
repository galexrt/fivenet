package completor

import (
	context "context"
	"database/sql"
	"errors"
	"strings"

	users "github.com/galexrt/fivenet/gen/go/proto/resources/users"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/mstlystcdata"
	"github.com/galexrt/fivenet/pkg/perms"
	"github.com/galexrt/fivenet/pkg/tracker"
	"github.com/galexrt/fivenet/pkg/utils"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"go.uber.org/fx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	tUsers    = table.Users.AS("usershort")
	tCategory = table.FivenetDocumentsCategories.AS("category")
)

var (
	ErrFailedSearch = status.Error(codes.Internal, "errors.CompletorService.ErrFailedSearch")
)

type Server struct {
	CompletorServiceServer

	db      *sql.DB
	p       perms.Permissions
	data    *mstlystcdata.Cache
	tracker *tracker.Tracker
}

type Params struct {
	fx.In

	DB      *sql.DB
	Perms   perms.Permissions
	Data    *mstlystcdata.Cache
	Tracker *tracker.Tracker
}

func NewServer(p Params) *Server {
	s := &Server{
		db:      p.DB,
		p:       p.Perms,
		data:    p.Data,
		tracker: p.Tracker,
	}

	return s
}

func (s *Server) CompleteCitizens(ctx context.Context, req *CompleteCitizensRequest) (*CompleteCitizensRespoonse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	req.Search = strings.TrimSpace(req.Search)
	req.Search = strings.ReplaceAll(req.Search, "%", "")
	req.Search = strings.ReplaceAll(req.Search, " ", "%")

	condition := jet.Bool(true)
	if req.Search != "" {
		req.Search = "%" + req.Search + "%"
		condition = jet.CONCAT(tUsers.Firstname, jet.String(" "), tUsers.Lastname).
			LIKE(jet.String(req.Search))
	}
	if req.CurrentJob != nil && *req.CurrentJob {
		condition = condition.AND(
			tUsers.Job.EQ(jet.String(userInfo.Job)),
		)
	}

	stmt := tUsers.
		SELECT(
			tUsers.ID,
			tUsers.Identifier,
			tUsers.Firstname,
			tUsers.Lastname,
		).
		OPTIMIZER_HINTS(jet.OptimizerHint("idx_users_firstname_lastname_fulltext")).
		FROM(tUsers).
		WHERE(condition).
		ORDER_BY(
			tUsers.Lastname.DESC(),
		).
		LIMIT(15)

	var dest []*users.UserShort
	if err := stmt.QueryContext(ctx, s.db, &dest); err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, ErrFailedSearch
		}
	}

	if req.OnDuty != nil && *req.OnDuty {
		for i := len(dest) - 1; i >= 0; i-- {
			if !s.tracker.IsUserOnDuty(userInfo.Job, dest[i].UserId) {
				dest = utils.RemoveFromSlice(dest, i)
			}
		}
	}

	return &CompleteCitizensRespoonse{
		Users: dest,
	}, nil
}

func (s *Server) CompleteJobs(ctx context.Context, req *CompleteJobsRequest) (*CompleteJobsResponse, error) {

	var search string
	if req.Search != nil {
		search = *req.Search
	}
	if req.CurrentJob != nil && *req.CurrentJob {
		userInfo := auth.MustGetUserInfoFromContext(ctx)
		search = userInfo.Job
	}
	exactMatch := false
	if req.ExactMatch != nil {
		exactMatch = *req.ExactMatch
	}

	resp := &CompleteJobsResponse{}
	var err error
	resp.Jobs, err = s.data.GetSearcher().SearchJobs(ctx, search, exactMatch)
	if err != nil {
		return nil, ErrFailedSearch
	}

	return resp, nil
}

func (s *Server) CompleteDocumentCategories(ctx context.Context, req *CompleteDocumentCategoriesRequest) (*CompleteDocumentCategoriesResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	jobsAttr, err := s.p.Attr(userInfo, CompletorServicePerm, CompletorServiceCompleteDocumentCategoriesPerm, CompletorServiceCompleteDocumentCategoriesJobsPermField)
	if err != nil {
		return nil, ErrFailedSearch
	}
	var jobs perms.StringList
	if jobsAttr != nil {
		jobs = jobsAttr.([]string)
	}

	if len(jobs) == 0 {
		jobs = append(jobs, userInfo.Job)
	}

	req.Search = strings.TrimSpace(req.Search)
	req.Search = strings.ReplaceAll(req.Search, "%", "")
	req.Search = strings.ReplaceAll(req.Search, " ", "%")

	jobsExp := make([]jet.Expression, len(jobs))
	for i := 0; i < len(jobs); i++ {
		jobsExp[i] = jet.String(jobs[i])
	}

	condition := tCategory.Job.IN(jobsExp...)
	if req.Search != "" {
		req.Search = "%" + req.Search + "%"
		condition = condition.AND(
			tCategory.Name.LIKE(jet.String(req.Search)),
		)
	}

	stmt := tCategory.
		SELECT(
			tCategory.ID,
			tCategory.Name,
			tCategory.Description,
			tCategory.Job,
		).
		FROM(tCategory).
		WHERE(condition).
		ORDER_BY(
			tCategory.Name.DESC(),
		).
		LIMIT(15)

	resp := &CompleteDocumentCategoriesResponse{}
	if err := stmt.QueryContext(ctx, s.db, &resp.Categories); err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, ErrFailedSearch
		}
	}

	return resp, nil
}

func (s *Server) ListLawBooks(ctx context.Context, req *ListLawBooksRequest) (*ListLawBooksResponse, error) {
	return &ListLawBooksResponse{
		Books: s.data.GetLawBooks(),
	}, nil
}
