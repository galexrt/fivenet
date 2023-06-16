package rector

import (
	"context"
	"errors"
	"fmt"

	"github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
	rector "github.com/galexrt/fivenet/gen/go/proto/resources/rector"
	"github.com/galexrt/fivenet/gen/go/proto/resources/timestamp"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/grpc/auth/userinfo"
	"github.com/galexrt/fivenet/pkg/perms"
	"github.com/galexrt/fivenet/pkg/perms/collections"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/go-jet/jet/v2/qrm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrFailedQuery       = status.Error(codes.Internal, "errors.RectorService.ErrFailedQuery")
	ErrInvalidRequest    = status.Error(codes.InvalidArgument, "errors.RectorService.ErrInvalidRequest")
	ErrNoPermission      = status.Error(codes.PermissionDenied, "errors.RectorService.ErrNoPermission")
	ErrRoleAlreadyExists = status.Error(codes.InvalidArgument, "errors.RectorService.ErrRoleAlreadyExists")
	ErrOwnRoleDeletion   = status.Error(codes.InvalidArgument, "errors.RectorService.ErrOwnRoleDeletion")
	ErrInvalidAttrs      = status.Error(codes.InvalidArgument, "errors.RectorService.ErrInvalidAttrs")
)

var (
	ignoredGuardPermissions = []string{}
)

func (s *Server) ensureUserCanAccessRole(ctx context.Context, roleId uint64) (*model.FivenetRoles, bool, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	role, err := s.p.GetRole(ctx, roleId)
	if err != nil {
		return nil, false, err
	}

	if userInfo.SuperUser {
		return role, true, nil
	}

	// Make sure the user is from the job
	if role.Job != userInfo.Job {
		return nil, false, ErrInvalidRequest
	}

	if role.Grade > userInfo.JobGrade {
		return nil, false, ErrInvalidRequest
	}

	return role, true, nil
}

func (s *Server) filterPermissions(ctx context.Context, ps []*permissions.Permission) ([]*permissions.Permission, error) {
	filtered := []*permissions.Permission{}

outer:
	for _, p := range ps {
		for i := 0; i < len(ignoredGuardPermissions); i++ {
			if p.GuardName == ignoredGuardPermissions[i] {
				continue outer
			}
		}

		filtered = append(filtered, p)
	}

	return filtered, nil
}

func (s *Server) filterPermissionIDs(ctx context.Context, ids []uint64) ([]uint64, error) {
	if len(ids) == 0 {
		return ids, nil
	}

	perms, err := s.p.GetPermissionsByIDs(ctx, ids...)
	if err != nil {
		return nil, err
	}

	filtered, err := s.filterPermissions(ctx, perms)
	if err != nil {
		return nil, err
	}

	permIds := make([]uint64, len(filtered))
	for i := 0; i < len(filtered); i++ {
		permIds[i] = filtered[i].Id
	}
	return permIds, nil
}

func (s *Server) filterAttributes(ctx context.Context, attrs []*permissions.RoleAttribute) error {
	if len(attrs) == 0 {
		return nil
	}

	for i := 0; i < len(attrs); i++ {
		attr, ok := s.p.GetRoleAttributeByID(attrs[i].RoleId, attrs[i].AttrId)
		if !ok {
			return fmt.Errorf("failed to find role %d attribute by id %d", attrs[i].RoleId, attrs[i].AttrId)
		}

		if !attrs[i].Value.Check(permissions.AttributeTypes(attr.Type), attr.ValidValues, attr.MaxValues) {
			return fmt.Errorf("failed to validate attribute values")
		}

		attrs[i].Type = attr.Type
	}

	return nil
}

func (s *Server) GetRoles(ctx context.Context, req *GetRolesRequest) (*GetRolesResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	var roles collections.Roles
	var err error

	if userInfo.SuperUser && req.All != nil && *req.All {
		roles, err = s.p.GetRoles(ctx, true)
		if err != nil {
			return nil, ErrFailedQuery
		}
	} else {
		roles, err = s.p.GetJobRolesUpTo(ctx, userInfo.Job, userInfo.JobGrade)
		if err != nil {
			return nil, ErrFailedQuery
		}
	}

	resp := &GetRolesResponse{}
	for _, r := range roles {
		role := &permissions.Role{
			Id:          r.ID,
			CreatedAt:   timestamp.New(*r.CreatedAt),
			Job:         r.Job,
			Grade:       r.Grade,
			Permissions: []*permissions.Permission{},
		}

		s.c.EnrichJobInfo(role)

		resp.Roles = append(resp.Roles, role)
	}

	return resp, nil
}

func (s *Server) GetRole(ctx context.Context, req *GetRoleRequest) (*GetRoleResponse, error) {
	role, check, err := s.ensureUserCanAccessRole(ctx, req.Id)
	if err != nil {
		return nil, ErrInvalidRequest
	}
	if !check {
		return nil, ErrNoPermission
	}

	perms, err := s.p.GetRolePermissions(ctx, role.ID)
	if err != nil {
		return nil, ErrInvalidRequest
	}

	resp := &GetRoleResponse{}

	resp.Role = &permissions.Role{
		Id:        role.ID,
		CreatedAt: timestamp.New(*role.CreatedAt),
		Job:       role.Job,
		Grade:     role.Grade,
	}
	s.c.EnrichJobInfo(resp.Role)

	fPerms, err := s.filterPermissions(ctx, perms)
	if err != nil {
		return nil, ErrInvalidRequest
	}

	resp.Role.Permissions = make([]*permissions.Permission, len(fPerms))
	copy(resp.Role.Permissions, fPerms)

	resp.Role.Attributes, err = s.p.GetRoleAttributes(role.Job, role.Grade)
	if err != nil {
		return nil, ErrFailedQuery
	}

	return resp, nil
}

func (s *Server) CreateRole(ctx context.Context, req *CreateRoleRequest) (*CreateRoleResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: RectorService_ServiceDesc.ServiceName,
		Method:  "CreateRole",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EVENT_TYPE_ERRORED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	role, err := s.p.GetRoleByJobAndGrade(ctx, userInfo.Job, req.Grade)
	if err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, ErrFailedQuery
		}
	}
	if role != nil {
		return nil, ErrRoleAlreadyExists
	}

	cr, err := s.p.CreateRole(ctx, userInfo.Job, req.Grade)
	if err != nil {
		return nil, ErrFailedQuery
	}

	if cr == nil {
		return nil, ErrInvalidRequest
	}

	auditEntry.State = int16(rector.EVENT_TYPE_CREATED)
	return &CreateRoleResponse{
		Role: permissions.ConvertFromRole(cr),
	}, nil
}

func (s *Server) DeleteRole(ctx context.Context, req *DeleteRoleRequest) (*DeleteRoleResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: RectorService_ServiceDesc.ServiceName,
		Method:  "DeleteRole",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EVENT_TYPE_ERRORED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	role, check, err := s.ensureUserCanAccessRole(ctx, req.Id)
	if err != nil {
		return nil, ErrInvalidRequest
	}
	if !check {
		return nil, ErrNoPermission
	}

	roleCount, err := s.p.CountRolesForJob(ctx, userInfo.Job)
	if err != nil {
		return nil, ErrInvalidRequest
	}

	// Don't allow deleting the "last" role, one role should always remain
	if roleCount <= 1 {
		return nil, ErrInvalidRequest
	}

	// Don't allow deleting the own or higher role
	if role.Grade >= userInfo.JobGrade {
		return nil, ErrOwnRoleDeletion
	}

	if err := s.p.DeleteRole(ctx, role.ID); err != nil {
		return nil, ErrInvalidRequest
	}

	auditEntry.State = int16(rector.EVENT_TYPE_DELETED)

	return &DeleteRoleResponse{}, nil
}

func (s *Server) UpdateRolePerms(ctx context.Context, req *UpdateRolePermsRequest) (*UpdateRolePermsResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: RectorService_ServiceDesc.ServiceName,
		Method:  "UpdateRolePerms",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EVENT_TYPE_ERRORED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	role, check, err := s.ensureUserCanAccessRole(ctx, req.Id)
	if err != nil {
		return nil, ErrInvalidRequest
	}
	if !check {
		return nil, ErrNoPermission
	}

	if req.Perms != nil {
		if err := s.handlPermissionsUpdate(ctx, role, req.Perms); err != nil {
			return nil, ErrInvalidAttrs
		}
	}
	if req.Attrs != nil {
		if err := s.handleAttributeUpdate(ctx, userInfo, role, req.Attrs); err != nil {
			return nil, ErrInvalidAttrs
		}
	}

	auditEntry.State = int16(rector.EVENT_TYPE_UPDATED)

	return &UpdateRolePermsResponse{}, nil
}

func (s *Server) handlPermissionsUpdate(ctx context.Context, role *model.FivenetRoles, permsUpdate *PermsUpdate) error {
	updatePermIds := make([]uint64, len(permsUpdate.ToUpdate))
	for i := 0; i < len(permsUpdate.ToUpdate); i++ {
		updatePermIds[i] = permsUpdate.ToUpdate[i].Id
	}
	toUpdate, err := s.filterPermissionIDs(ctx, updatePermIds)
	if err != nil {
		return err
	}

	toDelete, err := s.filterPermissionIDs(ctx, permsUpdate.ToRemove)
	if err != nil {
		return err
	}

	if len(toUpdate) > 0 {
		toUpdatePerms := make([]perms.AddPerm, len(permsUpdate.ToUpdate))
		for _, v := range toUpdate {
			for i := 0; i < len(permsUpdate.ToUpdate); i++ {
				if v == permsUpdate.ToUpdate[i].Id {
					toUpdatePerms[i] = perms.AddPerm{
						Id:  permsUpdate.ToUpdate[i].Id,
						Val: permsUpdate.ToUpdate[i].Val,
					}
					break
				}
			}
		}
		if err := s.p.UpdateRolePermissions(ctx, role.ID, toUpdatePerms...); err != nil {
			return err
		}
	}
	if len(toDelete) > 0 {
		if err := s.p.RemovePermissionsFromRole(ctx, role.ID, toDelete...); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) handleAttributeUpdate(ctx context.Context, userInfo *userinfo.UserInfo, role *model.FivenetRoles, attrUpdates *AttrsUpdate) error {
	if err := s.filterAttributes(ctx, attrUpdates.ToUpdate); err != nil {
		return err
	}

	if err := s.filterAttributes(ctx, attrUpdates.ToRemove); err != nil {
		return err
	}

	if len(attrUpdates.ToUpdate) > 0 {
		if err := s.p.AddOrUpdateAttributesToRole(ctx, userInfo.Job, userInfo.JobGrade, role.ID, attrUpdates.ToUpdate...); err != nil {
			return err
		}
	}
	if len(attrUpdates.ToRemove) > 0 {
		if err := s.p.RemoveAttributesFromRole(ctx, role.ID, attrUpdates.ToRemove...); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) GetPermissions(ctx context.Context, req *GetPermissionsRequest) (*GetPermissionsResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	perms, err := s.p.GetAllPermissions(ctx)
	if err != nil {
		return nil, ErrFailedQuery
	}

	filtered, err := s.filterPermissions(ctx, perms)
	if err != nil {
		return nil, ErrInvalidRequest
	}

	resp := &GetPermissionsResponse{}
	resp.Permissions = filtered

	role, err := s.p.GetRole(ctx, req.RoleId)
	if err != nil {
		return nil, ErrInvalidRequest
	}

	if role.Job != userInfo.Job && !userInfo.SuperUser {
		return nil, ErrInvalidRequest
	}

	attrs, err := s.p.GetAllAttributes(ctx, userInfo.Job, userInfo.JobGrade)
	if err != nil {
		return nil, ErrInvalidRequest
	}
	resp.Attributes = attrs

	return resp, nil
}
