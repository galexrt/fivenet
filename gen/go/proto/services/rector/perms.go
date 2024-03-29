package rector

import (
	"context"
	"errors"

	"github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
	rector "github.com/galexrt/fivenet/gen/go/proto/resources/rector"
	"github.com/galexrt/fivenet/gen/go/proto/resources/timestamp"
	errorsrector "github.com/galexrt/fivenet/gen/go/proto/services/rector/errors"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/grpc/auth/userinfo"
	"github.com/galexrt/fivenet/pkg/grpc/errswrap"
	"github.com/galexrt/fivenet/pkg/perms"
	"github.com/galexrt/fivenet/pkg/perms/collections"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/go-jet/jet/v2/qrm"
)

var (
	ignoredGuardPermissions = []string{}
)

func (s *Server) ensureUserCanAccessRole(ctx context.Context, roleId uint64) (*model.FivenetRoles, bool, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	role, err := s.ps.GetRole(ctx, roleId)
	if err != nil {
		return nil, false, err
	}

	if userInfo.SuperUser {
		return role, true, nil
	}

	// Make sure the user is from the job
	if role.Job != userInfo.Job {
		return nil, false, errorsrector.ErrInvalidRequest
	}

	if role.Grade > userInfo.JobGrade {
		return nil, false, errorsrector.ErrInvalidRequest
	}

	return role, true, nil
}

func (s *Server) filterPermissions(ctx context.Context, job string, fullFilter bool, ps []*permissions.Permission) ([]*permissions.Permission, error) {
	filtered := []*permissions.Permission{}

	filters, err := s.ps.GetJobPermissions(ctx, job)
	if err != nil {
		return nil, err
	}

outer:
	for _, p := range ps {
		for i := 0; i < len(ignoredGuardPermissions); i++ {
			if p.GuardName == ignoredGuardPermissions[i] {
				continue outer
			}
		}

		if fullFilter {
			found := false
			for _, filter := range filters {
				if p.Id == filter.Id {
					if !filter.Val {
						continue outer
					}
					found = true
				}
			}
			if !found {
				continue
			}
		}

		filtered = append(filtered, p)
	}

	return filtered, nil
}

func (s *Server) filterPermissionIDs(ctx context.Context, job string, full bool, ids []uint64) ([]uint64, error) {
	if len(ids) == 0 {
		return ids, nil
	}

	perms, err := s.ps.GetPermissionsByIDs(ctx, ids...)
	if err != nil {
		return nil, err
	}

	filtered, err := s.filterPermissions(ctx, job, full, perms)
	if err != nil {
		return nil, err
	}

	permIds := make([]uint64, len(filtered))
	for i := 0; i < len(filtered); i++ {
		permIds[i] = filtered[i].Id
	}
	return permIds, nil
}

func (s *Server) GetRoles(ctx context.Context, req *GetRolesRequest) (*GetRolesResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	var roles collections.Roles
	var err error

	if userInfo.SuperUser && req.LowestRank != nil && *req.LowestRank {
		roles, err = s.ps.GetRoles(ctx, true)
		if err != nil {
			return nil, errswrap.NewError(errorsrector.ErrFailedQuery, err)
		}

		collectedRoles := map[string]*model.FivenetRoles{}
		for _, role := range roles {
			if _, ok := collectedRoles[role.Job]; !ok {
				collectedRoles[role.Job] = role
				continue
			}
		}

		roles = collections.Roles{}
		for _, role := range collectedRoles {
			roles = append(roles, role)
		}
	} else {
		roles, err = s.ps.GetJobRolesUpTo(ctx, userInfo.Job, userInfo.JobGrade)
		if err != nil {
			return nil, errswrap.NewError(errorsrector.ErrFailedQuery, err)
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

		s.enricher.EnrichJobInfo(role)

		resp.Roles = append(resp.Roles, role)
	}

	return resp, nil
}

func (s *Server) GetRole(ctx context.Context, req *GetRoleRequest) (*GetRoleResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	role, check, err := s.ensureUserCanAccessRole(ctx, req.Id)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}
	if !check && !userInfo.SuperUser {
		return nil, errorsrector.ErrNoPermission
	}

	fullFilter := !(userInfo.SuperUser && req.Filtered != nil && !*req.Filtered)

	var perms []*permissions.Permission
	if fullFilter {
		perms, err = s.ps.GetRolePermissions(ctx, role.ID)
		if err != nil {
			return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
		}
	} else {
		perms, err = s.ps.GetJobPermissions(ctx, role.Job)
		if err != nil {
			return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
		}
	}

	resp := &GetRoleResponse{
		Role: &permissions.Role{
			Id:        role.ID,
			CreatedAt: timestamp.New(*role.CreatedAt),
			Job:       role.Job,
			Grade:     role.Grade,
		},
	}
	s.enricher.EnrichJobInfo(resp.Role)

	fPerms, err := s.filterPermissions(ctx, role.Job, fullFilter, perms)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}

	resp.Role.Permissions = make([]*permissions.Permission, len(fPerms))
	copy(resp.Role.Permissions, fPerms)

	resp.Role.Attributes, err = s.ps.GetRoleAttributes(role.Job, role.Grade)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrFailedQuery, err)
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
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	// Make sure the user is from the job or is a super user
	if !userInfo.SuperUser {
		if req.Job != userInfo.Job {
			return nil, errorsrector.ErrInvalidRequest
		}
		if req.Grade > userInfo.JobGrade {
			return nil, errorsrector.ErrInvalidRequest
		}
	}

	role, err := s.ps.GetRoleByJobAndGrade(ctx, req.Job, req.Grade)
	if err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, errswrap.NewError(errorsrector.ErrFailedQuery, err)
		}
	}
	if role != nil {
		return nil, errorsrector.ErrRoleAlreadyExists
	}

	cr, err := s.ps.CreateRole(ctx, req.Job, req.Grade)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrFailedQuery, err)
	}

	if cr == nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}

	r := permissions.ConvertFromRole(cr)
	s.enricher.EnrichJobInfo(r)

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_CREATED)
	return &CreateRoleResponse{
		Role: r,
	}, nil
}

func (s *Server) DeleteRole(ctx context.Context, req *DeleteRoleRequest) (*DeleteRoleResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: RectorService_ServiceDesc.ServiceName,
		Method:  "DeleteRole",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	role, check, err := s.ensureUserCanAccessRole(ctx, req.Id)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}
	if !check && !userInfo.SuperUser {
		return nil, errorsrector.ErrNoPermission
	}

	roleCount, err := s.ps.CountRolesForJob(ctx, userInfo.Job)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}

	// Don't allow deleting the "last" role, one role should always remain
	if roleCount <= 1 {
		return nil, errorsrector.ErrInvalidRequest
	}

	// Don't allow deleting the own or higher role
	if role.Grade >= userInfo.JobGrade {
		return nil, errorsrector.ErrOwnRoleDeletion
	}

	if err := s.ps.DeleteRole(ctx, role.ID); err != nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_DELETED)

	return &DeleteRoleResponse{}, nil
}

func (s *Server) UpdateRolePerms(ctx context.Context, req *UpdateRolePermsRequest) (*UpdateRolePermsResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: RectorService_ServiceDesc.ServiceName,
		Method:  "UpdateRolePerms",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.aud.Log(auditEntry, req)

	role, check, err := s.ensureUserCanAccessRole(ctx, req.Id)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}
	if !check && !userInfo.SuperUser {
		return nil, errswrap.NewError(errorsrector.ErrNoPermission, err)
	}

	if req.Perms != nil {
		if err := s.handlPermissionsUpdate(ctx, role, req.Perms); err != nil {
			return nil, errswrap.NewError(errorsrector.ErrInvalidPerms, err)
		}
	}
	if req.Attrs != nil {
		if err := s.handleAttributeUpdate(ctx, userInfo, role, req.Attrs); err != nil {
			return nil, errswrap.NewError(errorsrector.ErrInvalidAttrs, err)
		}
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	return &UpdateRolePermsResponse{}, nil
}

func (s *Server) handlPermissionsUpdate(ctx context.Context, role *model.FivenetRoles, permsUpdate *PermsUpdate) error {
	updatePermIds := make([]uint64, len(permsUpdate.ToUpdate))
	for i := 0; i < len(permsUpdate.ToUpdate); i++ {
		updatePermIds[i] = permsUpdate.ToUpdate[i].Id
	}
	toUpdate, err := s.filterPermissionIDs(ctx, role.Job, false, updatePermIds)
	if err != nil {
		return err
	}

	toDelete, err := s.filterPermissionIDs(ctx, role.Job, false, permsUpdate.ToRemove)
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
		if err := s.ps.UpdateRolePermissions(ctx, role.ID, toUpdatePerms...); err != nil {
			return err
		}
	}
	if len(toDelete) > 0 {
		if err := s.ps.RemovePermissionsFromRole(ctx, role.ID, toDelete...); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) handleAttributeUpdate(ctx context.Context, userInfo *userinfo.UserInfo, role *model.FivenetRoles, attrUpdates *AttrsUpdate) error {
	if len(attrUpdates.ToUpdate) > 0 {
		if err := s.ps.AddOrUpdateAttributesToRole(ctx, userInfo.Job, role.ID, attrUpdates.ToUpdate...); err != nil {
			return err
		}
	}

	if len(attrUpdates.ToRemove) > 0 {
		if err := s.ps.RemoveAttributesFromRole(ctx, role.ID, attrUpdates.ToRemove...); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) GetPermissions(ctx context.Context, req *GetPermissionsRequest) (*GetPermissionsResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	perms, err := s.ps.GetAllPermissions(ctx)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrFailedQuery, err)
	}

	fullFilter := !(userInfo.SuperUser && req.Filtered != nil && !*req.Filtered)

	filtered, err := s.filterPermissions(ctx, userInfo.Job, fullFilter, perms)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}

	resp := &GetPermissionsResponse{}
	resp.Permissions = filtered

	role, err := s.ps.GetRole(ctx, req.RoleId)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}

	if role.Job != userInfo.Job && !userInfo.SuperUser {
		return nil, errorsrector.ErrInvalidRequest
	}

	attrs, err := s.ps.GetAllAttributes(ctx, role.Job, role.Grade)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}
	resp.Attributes = attrs

	return resp, nil
}

func (s *Server) UpdateRoleLimits(ctx context.Context, req *UpdateRoleLimitsRequest) (*UpdateRoleLimitsResponse, error) {
	role, err := s.ps.GetRole(ctx, req.RoleId)
	if err != nil {
		return nil, errswrap.NewError(errorsrector.ErrInvalidRequest, err)
	}
	for _, attr := range req.Attrs.ToUpdate {
		if err := s.ps.UpdateJobAttributeMaxValues(ctx, role.Job, attr.AttrId, attr.MaxValues); err != nil {
			return nil, errswrap.NewError(errorsrector.ErrFailedQuery, err)
		}
	}

	for _, ps := range req.Perms.ToUpdate {
		if err := s.ps.UpdateJobPermissions(ctx, role.Job, ps.Id, ps.Val); err != nil {
			return nil, errswrap.NewError(errorsrector.ErrFailedQuery, err)
		}
	}

	for _, ps := range req.Perms.ToRemove {
		if err := s.ps.UpdateJobPermissions(ctx, role.Job, ps, false); err != nil {
			return nil, errswrap.NewError(errorsrector.ErrFailedQuery, err)
		}
	}

	if err := s.ps.ApplyJobPermissions(ctx, role.Job); err != nil {
		return nil, errswrap.NewError(errorsrector.ErrFailedQuery, err)
	}

	return &UpdateRoleLimitsResponse{}, nil
}
