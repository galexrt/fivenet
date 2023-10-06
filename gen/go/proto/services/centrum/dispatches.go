package centrum

import (
	"context"

	database "github.com/galexrt/fivenet/gen/go/proto/resources/common/database"
	dispatch "github.com/galexrt/fivenet/gen/go/proto/resources/dispatch"
	"github.com/galexrt/fivenet/gen/go/proto/resources/rector"
	"github.com/galexrt/fivenet/gen/go/proto/resources/timestamp"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/utils"
	"github.com/galexrt/fivenet/pkg/utils/dbutils"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var (
	tDispatch       = table.FivenetCentrumDispatches.AS("dispatch")
	tDispatchStatus = table.FivenetCentrumDispatchesStatus.AS("dispatchstatus")
	tDispatchUnit   = table.FivenetCentrumDispatchesAsgmts.AS("dispatchassignment")
)

var (
	ErrModeForbidsAction        = status.Error(codes.InvalidArgument, "errors.CentrumService.ErrModeForbidsAction.title;errors.CentrumService.ErrModeForbidsAction.content")
	ErrDispatchAlreadyCompleted = status.Error(codes.InvalidArgument, "errors.CentrumService.ErrDispatchAlreadyCompleted.title;errors.CentrumService.ErrDispatchAlreadyCompleted.content")
)

func (s *Server) ListDispatches(ctx context.Context, req *ListDispatchesRequest) (*ListDispatchesResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "ListDispatches",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	resp := &ListDispatchesResponse{
		Dispatches: []*dispatch.Dispatch{},
	}

	dispatches := s.listDispatches(userInfo.Job)
	for i := 0; i < len(dispatches); i++ {
		// Hide user info when dispatch is anonymous
		if dispatches[i].Anon {
			dispatches[i].Creator = nil
			dispatches[i].CreatorId = nil
		}

		include := true

		// Include statuses that should be listed
		if len(req.Status) > 0 && !utils.InSlice(req.Status, dispatches[i].Status.Status) {
			include = false
		} else if len(req.NotStatus) > 0 {
			// Which statuses to ignore
			for _, status := range req.NotStatus {
				if dispatches[i].Status != nil && dispatches[i].Status.Status == status {
					include = false
					break
				}
			}
		}

		if include {
			resp.Dispatches = append(resp.Dispatches, dispatches[i])
		}
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_VIEWED)

	return resp, nil
}

func (s *Server) CreateDispatch(ctx context.Context, req *CreateDispatchRequest) (*CreateDispatchResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "CreateDispatch",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	req.Dispatch.Job = userInfo.Job
	req.Dispatch.CreatorId = &userInfo.UserId

	dsp, err := s.createDispatch(ctx, req.Dispatch)
	if err != nil {
		return nil, err
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_CREATED)

	return &CreateDispatchResponse{
		Dispatch: dsp,
	}, nil
}

func (s *Server) createDispatch(ctx context.Context, d *dispatch.Dispatch) (*dispatch.Dispatch, error) {
	postal := s.postals.Closest(d.X, d.Y)
	if postal != nil {
		d.Postal = postal.Code
	}

	// Begin transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Defer a rollback in case anything fails
	defer tx.Rollback()

	tDispatch := table.FivenetCentrumDispatches
	stmt := tDispatch.
		INSERT(
			tDispatch.Job,
			tDispatch.Message,
			tDispatch.Description,
			tDispatch.Attributes,
			tDispatch.X,
			tDispatch.Y,
			tDispatch.Postal,
			tDispatch.Anon,
			tDispatch.CreatorID,
		).
		VALUES(
			d.Job,
			d.Message,
			d.Description,
			d.Attributes,
			d.X,
			d.Y,
			d.Postal,
			d.Anon,
			d.CreatorId,
		)

	result, err := stmt.ExecContext(ctx, tx)
	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	if err := s.addDispatchStatus(ctx, tx, &dispatch.DispatchStatus{
		DispatchId: uint64(lastId),
		UserId:     d.CreatorId,
		Status:     dispatch.StatusDispatch_STATUS_DISPATCH_NEW,
		X:          &d.X,
		Y:          &d.Y,
		Postal:     d.Postal,
	}); err != nil {
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// Load dispatch into cache
	if err := s.loadDispatches(ctx, uint64(lastId)); err != nil {
		return nil, err
	}

	dsp, ok := s.getDispatch(d.Job, uint64(lastId))
	if !ok {
		return nil, ErrFailedQuery
	}
	// Hide user info when dispatch is anonymous
	if dsp.Anon {
		dsp.Creator = nil
		dsp.CreatorId = nil
	}

	data, err := proto.Marshal(dsp)
	if err != nil {
		return nil, err
	}
	s.events.JS.PublishAsync(buildSubject(TopicDispatch, TypeDispatchCreated, d.Job, 0), data)

	return dsp, nil
}

func (s *Server) UpdateDispatch(ctx context.Context, req *UpdateDispatchRequest) (*UpdateDispatchResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "UpdateDispatch",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	if err := s.updateDispatch(ctx, userInfo, req.Dispatch); err != nil {
		return nil, ErrFailedQuery
	}

	// Load dispatch into cache
	if err := s.loadDispatches(ctx, req.Dispatch.Id); err != nil {
		return nil, ErrFailedQuery
	}

	dsp, ok := s.getDispatch(userInfo.Job, req.Dispatch.Id)
	if !ok {
		return nil, ErrFailedQuery
	}

	data, err := proto.Marshal(dsp)
	if err != nil {
		return nil, ErrFailedQuery
	}
	s.events.JS.PublishAsync(buildSubject(TopicDispatch, TypeDispatchUpdated, userInfo.Job, 0), data)
	for _, unit := range dsp.Units {
		s.events.JS.PublishAsync(buildSubject(TopicDispatch, TypeDispatchUpdated, userInfo.Job, unit.UnitId), data)
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	return &UpdateDispatchResponse{}, nil
}

func (s *Server) TakeDispatch(ctx context.Context, req *TakeDispatchRequest) (*TakeDispatchResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "TakeDispatch",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	unitId, ok := s.getUnitIDForUserID(userInfo.UserId)
	if !ok {
		return nil, ErrFailedQuery
	}
	unit, ok := s.getUnit(userInfo.Job, unitId)
	if !ok {
		return nil, ErrFailedQuery
	}

	settings := s.getSettings(userInfo.Job)
	for _, dispatchId := range req.DispatchIds {
		dsp, ok := s.getDispatch(userInfo.Job, dispatchId)
		if !ok {
			return nil, ErrFailedQuery
		}

		// If the dispatch center is in central command mode, units can't self assign dispatches
		if settings.Mode == dispatch.CentrumMode_CENTRUM_MODE_CENTRAL_COMMAND {
			if !utils.InSliceFunc(dsp.Units, func(in *dispatch.DispatchAssignment) bool {
				return in.UnitId == unitId
			}) {
				return nil, ErrModeForbidsAction
			}
		}

		// If dispatch is completed, disallow to accept the dispatch
		if dsp.Status != nil && s.isStatusDispatchComplete(dsp.Status.Status) {
			return nil, ErrDispatchAlreadyCompleted
		}

		var x, y *float64
		var postal *string
		marker, ok := s.tracker.GetUserById(userInfo.UserId)
		if ok {
			x = &marker.Info.X
			y = &marker.Info.Y
			postal = marker.Info.Postal
		}

		status := dispatch.StatusDispatch_STATUS_DISPATCH_UNSPECIFIED

		tDispatchUnit := table.FivenetCentrumDispatchesAsgmts
		// Dispatch accepted
		if req.Resp == TakeDispatchResp_TAKE_DISPATCH_RESP_ACCEPTED {
			status = dispatch.StatusDispatch_STATUS_DISPATCH_UNIT_ACCEPTED

			stmt := tDispatchUnit.
				INSERT(
					tDispatchUnit.DispatchID,
					tDispatchUnit.UnitID,
					tDispatchUnit.ExpiresAt,
				).
				VALUES(
					dsp.Id,
					unit.Id,
					jet.NULL,
				).
				ON_DUPLICATE_KEY_UPDATE(
					tDispatchUnit.ExpiresAt.SET(jet.TimestampExp(jet.NULL)),
				)

			if _, err := stmt.ExecContext(ctx, s.db); err != nil {
				if !dbutils.IsDuplicateError(err) {
					return nil, err
				}
			}

			found := false
			accepted := true
			// Set unit expires at to nil
			for _, ua := range dsp.Units {
				if ua.UnitId == unit.Id {
					found = true
					// If there's no expiration time the unit has been directly assigned
					if ua.ExpiresAt == nil {
						accepted = false
					}
					ua.ExpiresAt = nil
					break
				}
			}

			if !found {
				dsp.Units = append(dsp.Units, &dispatch.DispatchAssignment{
					DispatchId: dsp.Id,
					UnitId:     unit.Id,
					Unit:       unit,
					CreatedAt:  timestamp.Now(),
				})
			}

			if accepted {
				// Set unit to busy when unit accepts a dispatch
				if unit.Status == nil || unit.Status.Status != dispatch.StatusUnit_STATUS_UNIT_BUSY {
					if err := s.updateUnitStatus(ctx, userInfo.Job, unit, &dispatch.UnitStatus{
						UnitId:    unit.Id,
						Status:    dispatch.StatusUnit_STATUS_UNIT_BUSY,
						UserId:    &userInfo.UserId,
						CreatorId: &userInfo.UserId,
						X:         x,
						Y:         y,
						Postal:    postal,
					}); err != nil {
						return nil, err
					}
				}
			}
		} else {
			// Dispatch declined
			status = dispatch.StatusDispatch_STATUS_DISPATCH_UNIT_DECLINED

			stmt := tDispatchUnit.
				DELETE().
				WHERE(jet.AND(
					tDispatchUnit.DispatchID.EQ(jet.Uint64(dsp.Id)),
					tDispatchUnit.UnitID.EQ(jet.Uint64(unit.Id)),
				)).
				LIMIT(1)

			if _, err := stmt.ExecContext(ctx, s.db); err != nil {
				if !dbutils.IsDuplicateError(err) {
					return nil, err
				}
			}

			// Remove the unit's assignment
			for k, u := range dsp.Units {
				if u.UnitId == unit.Id {
					dsp.Units = utils.RemoveFromSlice(dsp.Units, k)
					break
				}
			}
		}

		if err := s.updateDispatchStatus(ctx, userInfo.Job, dsp, &dispatch.DispatchStatus{
			DispatchId: dispatchId,
			Status:     status,
			UnitId:     &unitId,
			UserId:     &userInfo.UserId,
			X:          x,
			Y:          y,
			Postal:     postal,
		}); err != nil {
			return nil, err
		}
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	return &TakeDispatchResponse{}, nil
}

func (s *Server) UpdateDispatchStatus(ctx context.Context, req *UpdateDispatchStatusRequest) (*UpdateDispatchStatusResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "UpdateDispatchStatus",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	dsp, ok := s.getDispatch(userInfo.Job, req.DispatchId)
	if !ok {
		return nil, ErrFailedQuery
	}

	if !s.checkIfUserIsPartOfDispatch(userInfo, dsp, false) && !userInfo.SuperUser {
		return nil, ErrNotPartOfDispatch
	}

	var statusUnitId *uint64
	unitId, ok := s.getUnitIDForUserID(userInfo.UserId)
	if !ok {
		if !s.checkIfUserIsDisponent(userInfo.Job, userInfo.UserId) {
			return nil, ErrFailedQuery
		}
	} else {
		statusUnitId = &unitId
	}

	var x, y *float64
	var postal *string
	marker, ok := s.tracker.GetUserById(userInfo.UserId)
	if ok {
		x = &marker.Info.X
		y = &marker.Info.Y
		postal = marker.Info.Postal
	}

	if err := s.updateDispatchStatus(ctx, userInfo.Job, dsp, &dispatch.DispatchStatus{
		DispatchId: dsp.Id,
		UnitId:     statusUnitId,
		Status:     req.Status,
		Code:       req.Code,
		Reason:     req.Reason,
		UserId:     &userInfo.UserId,
		X:          x,
		Y:          y,
		Postal:     postal,
	}); err != nil {
		return nil, err
	}

	if req.Status == dispatch.StatusDispatch_STATUS_DISPATCH_EN_ROUTE ||
		req.Status == dispatch.StatusDispatch_STATUS_DISPATCH_ON_SCENE ||
		req.Status == dispatch.StatusDispatch_STATUS_DISPATCH_NEED_ASSISTANCE {
		unit, ok := s.getUnit(userInfo.Job, unitId)
		if ok && unit != nil {
			// Set unit to busy when unit accepts a dispatch
			if unit.Status == nil || unit.Status.Status != dispatch.StatusUnit_STATUS_UNIT_BUSY {
				if err := s.updateUnitStatus(ctx, userInfo.Job, unit, &dispatch.UnitStatus{
					UnitId:    unit.Id,
					Status:    dispatch.StatusUnit_STATUS_UNIT_BUSY,
					UserId:    &userInfo.UserId,
					CreatorId: &userInfo.UserId,
					X:         x,
					Y:         y,
					Postal:    postal,
				}); err != nil {
					return nil, err
				}
			}
		}
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	return &UpdateDispatchStatusResponse{}, nil
}

func (s *Server) AssignDispatch(ctx context.Context, req *AssignDispatchRequest) (*AssignDispatchResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "AssignDispatch",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	dsp, ok := s.getDispatch(userInfo.Job, req.DispatchId)
	if !ok {
		return nil, ErrFailedQuery
	}

	if dsp.Job != userInfo.Job {
		return nil, ErrFailedQuery
	}

	if err := s.updateDispatchAssignments(ctx, userInfo.Job, &userInfo.UserId, dsp, req.ToAdd, req.ToRemove); err != nil {
		return nil, err
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	return &AssignDispatchResponse{}, nil
}

func (s *Server) ListDispatchActivity(ctx context.Context, req *ListDispatchActivityRequest) (*ListDispatchActivityResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	countStmt := tDispatchStatus.
		SELECT(
			jet.COUNT(jet.DISTINCT(tDispatchStatus.ID)).AS("datacount.totalcount"),
		).
		FROM(tDispatchStatus).
		WHERE(
			tDispatchStatus.DispatchID.EQ(jet.Uint64(req.Id)),
		)

	var count database.DataCount
	if err := countStmt.QueryContext(ctx, s.db, &count); err != nil {
		return nil, ErrFailedQuery
	}

	pag, limit := req.Pagination.GetResponseWithPageSize(10)
	resp := &ListDispatchActivityResponse{
		Pagination: pag,
	}
	if count.TotalCount <= 0 {
		return resp, nil
	}

	stmt := tDispatchStatus.
		SELECT(
			tDispatchStatus.ID,
			tDispatchStatus.CreatedAt,
			tDispatchStatus.UnitID,
			tDispatchStatus.Status,
			tDispatchStatus.Reason,
			tDispatchStatus.Code,
			tDispatchStatus.UserID,
			tDispatchStatus.X,
			tDispatchStatus.Y,
			tDispatchStatus.Postal,
			tUsers.Firstname,
			tUsers.Lastname,
			tUsers.Job,
			tUsers.JobGrade,
			tUsers.Dateofbirth,
			tUsers.PhoneNumber,
		).
		FROM(
			tDispatchStatus.
				LEFT_JOIN(tUsers,
					tUsers.ID.EQ(tDispatchStatus.UserID),
				),
		).
		WHERE(
			tDispatchStatus.DispatchID.EQ(jet.Uint64(req.Id)),
		).
		ORDER_BY(tDispatchStatus.ID.DESC()).
		OFFSET(req.Pagination.Offset).
		LIMIT(limit)

	if err := stmt.QueryContext(ctx, s.db, &resp.Activity); err != nil {
		return nil, err
	}
	for i := 0; i < len(resp.Activity); i++ {
		if resp.Activity[i].UnitId != nil && *resp.Activity[i].UnitId > 0 {
			resp.Activity[i].Unit, _ = s.getUnit(userInfo.Job, *resp.Activity[i].UnitId)
		}
	}

	resp.Pagination.Update(count.TotalCount, len(resp.Activity))

	return resp, nil
}

func (s *Server) DeleteDispatch(ctx context.Context, req *DeleteDispatchRequest) (*DeleteDispatchResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "DeleteDispatch",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	dsp, ok := s.getDispatch(userInfo.Job, req.Id)
	if !ok {
		return nil, ErrFailedQuery
	}

	stmt := tDispatch.
		DELETE().
		WHERE(jet.AND(
			tDispatch.Job.EQ(jet.String(userInfo.Job)),
			tDispatch.ID.EQ(jet.Uint64(dsp.Id)),
		)).
		LIMIT(1)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, ErrFailedQuery
	}

	data, err := proto.Marshal(dsp)
	if err != nil {
		return nil, err
	}
	s.events.JS.PublishAsync(buildSubject(TopicDispatch, TypeDispatchDeleted, dsp.Job, 0), data)

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_DELETED)

	return &DeleteDispatchResponse{}, nil
}
