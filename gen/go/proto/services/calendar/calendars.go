package calendar

import (
	"context"
	"errors"

	calendar "github.com/fivenet-app/fivenet/gen/go/proto/resources/calendar"
	database "github.com/fivenet-app/fivenet/gen/go/proto/resources/common/database"
	errorscalendar "github.com/fivenet-app/fivenet/gen/go/proto/services/calendar/errors"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth/userinfo"
	"github.com/fivenet-app/fivenet/pkg/grpc/errswrap"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
)

func (s *Server) ListCalendars(ctx context.Context, req *ListCalendarsRequest) (*ListCalendarsResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	condition := jet.AND(
		tCalendar.DeletedAt.IS_NULL(),
		jet.OR(
			jet.OR(
				tCalendar.Public.IS_TRUE(),
				tCalendar.CreatorID.EQ(jet.Int32(userInfo.UserId)),
			),
			jet.OR(
				jet.AND(
					tCUserAccess.Access.IS_NOT_NULL(),
					tCUserAccess.Access.NOT_EQ(jet.Int32(int32(calendar.AccessLevel_ACCESS_LEVEL_BLOCKED))),
				),
				jet.AND(
					tCUserAccess.Access.IS_NULL(),
					tCJobAccess.Access.IS_NOT_NULL(),
					tCJobAccess.Access.NOT_EQ(jet.Int32(int32(calendar.AccessLevel_ACCESS_LEVEL_BLOCKED))),
				),
			),
		),
	)

	countStmt := tCalendar.
		SELECT(
			jet.COUNT(jet.DISTINCT(tCalendar.ID)).AS("datacount.totalcount"),
		).
		FROM(tCalendar.
			LEFT_JOIN(tCUserAccess,
				tCUserAccess.CalendarID.EQ(tCalendar.ID).
					AND(tCUserAccess.EntryID.IS_NULL()).
					AND(tCUserAccess.UserID.EQ(jet.Int32(userInfo.UserId))),
			).
			LEFT_JOIN(tCJobAccess,
				tCJobAccess.CalendarID.EQ(tCalendar.ID).
					AND(tCJobAccess.EntryID.IS_NULL()).
					AND(tCJobAccess.Job.EQ(jet.String(userInfo.Job))).
					AND(tCJobAccess.MinimumGrade.LT_EQ(jet.Int32(userInfo.JobGrade))),
			).
			LEFT_JOIN(tCreator,
				tCalendar.CreatorID.EQ(tCreator.ID),
			),
		).
		GROUP_BY(tCalendar.ID).
		WHERE(condition)

	var count database.DataCount
	if err := countStmt.QueryContext(ctx, s.db, &count); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
		}
	}

	pag, limit := req.Pagination.GetResponse(count.TotalCount)
	resp := &ListCalendarsResponse{
		Pagination: pag,
	}

	if count.TotalCount <= 0 {
		return resp, nil
	}

	stmt := tCalendar.
		SELECT(
			tCalendar.ID,
			tCalendar.CreatedAt,
			tCalendar.UpdatedAt,
			tCalendar.DeletedAt,
			tCalendar.Job,
			tCalendar.Name,
			tCalendar.Description,
			tCalendar.Public,
			tCalendar.Closed,
			tCalendar.CreatorID,
			tCalendar.CreatorJob,
		).
		FROM(tCalendar.
			LEFT_JOIN(tCUserAccess,
				tCUserAccess.CalendarID.EQ(tCalendar.ID).
					AND(tCUserAccess.EntryID.IS_NULL()).
					AND(tCUserAccess.UserID.EQ(jet.Int32(userInfo.UserId))),
			).
			LEFT_JOIN(tCJobAccess,
				tCJobAccess.CalendarID.EQ(tCalendar.ID).
					AND(tCJobAccess.EntryID.IS_NULL()).
					AND(tCJobAccess.Job.EQ(jet.String(userInfo.Job))).
					AND(tCJobAccess.MinimumGrade.LT_EQ(jet.Int32(userInfo.JobGrade))),
			).
			LEFT_JOIN(tCreator,
				tCalendar.CreatorID.EQ(tCreator.ID),
			),
		).
		GROUP_BY(tCalendar.ID).
		WHERE(condition).
		OFFSET(req.Pagination.Offset).
		LIMIT(limit)

	if err := stmt.QueryContext(ctx, s.db, &resp.Calendars); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
		}
	}

	jobInfoFn := s.enricher.EnrichJobInfoSafeFunc(userInfo)
	for i := 0; i < len(resp.Calendars); i++ {
		if resp.Calendars[i].Creator != nil {
			jobInfoFn(resp.Calendars[i].Creator)
		}
	}

	resp.Pagination.Update(len(resp.Calendars))

	return resp, nil
}

func (s *Server) GetCalendar(ctx context.Context, req *GetCalendarRequest) (*GetCalendarResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	// Check if user has access to existing calendar
	check, err := s.checkIfUserHasAccessToCalendar(ctx, req.CalendarId, userInfo, calendar.AccessLevel_ACCESS_LEVEL_VIEW)
	if err != nil {
		return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
	}
	if !check {
		return nil, errswrap.NewError(err, errorscalendar.ErrNoPerms)
	}

	calendar, err := s.getCalendar(ctx, userInfo, tCalendar.ID.EQ(jet.Uint64(req.CalendarId)))
	if err != nil {
		return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
	}
	if calendar == nil {
		return nil, errswrap.NewError(err, errorscalendar.ErrNoPerms)
	}

	return &GetCalendarResponse{
		Calendar: calendar,
	}, nil
}

func (s *Server) CreateOrUpdateCalendar(ctx context.Context, req *CreateOrUpdateCalendarRequest) (*CreateOrUpdateCalendarResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	// Check if user has access to existing calendar
	if req.Calendar.Id > 0 {
		check, err := s.checkIfUserHasAccessToCalendar(ctx, req.Calendar.Id, userInfo, calendar.AccessLevel_ACCESS_LEVEL_MANAGE)
		if err != nil {
			return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
		}
		if !check {
			return nil, errswrap.NewError(err, errorscalendar.ErrNoPerms)
		}

		if req.Calendar.Description == nil {
			empty := ""
			req.Calendar.Description = &empty
		}

		stmt := tCalendar.
			UPDATE(
				tCalendar.Name,
				tCalendar.Description,
				tCalendar.Public,
				tCalendar.Closed,
			).
			SET(
				tCalendar.Name.SET(jet.String(req.Calendar.Name)),
				tCalendar.Description.SET(jet.String(*req.Calendar.Description)),
				tCalendar.Public.SET(jet.Bool(req.Calendar.Public)),
				tCalendar.Closed.SET(jet.Bool(req.Calendar.Closed)),
			)

		if _, err := stmt.ExecContext(ctx, s.db); err != nil {
			return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
		}
	} else {
		stmt := tCalendar.
			INSERT(
				tCalendar.Job,
				tCalendar.Name,
				tCalendar.Description,
				tCalendar.Public,
				tCalendar.Closed,
				tCalendar.CreatorID,
				tCalendar.CreatorJob,
			).
			VALUES(
				req.Calendar.Job,
				req.Calendar.Name,
				req.Calendar.Description,
				req.Calendar.Public,
				req.Calendar.Closed,
				userInfo.UserId,
				userInfo.Job,
			).
			ON_DUPLICATE_KEY_UPDATE(
				tCalendar.Name.SET(jet.String(req.Calendar.Name)),
				tCalendar.Description.SET(jet.String("VALUES(`description`)")),
				tCalendar.Public.SET(jet.Bool(req.Calendar.Public)),
				tCalendar.Closed.SET(jet.Bool(req.Calendar.Closed)),
			)
		res, err := stmt.ExecContext(ctx, s.db)
		if err != nil {
			return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
		}

		if req.Calendar.Id == 0 {
			lastId, err := res.LastInsertId()
			if err != nil {
				return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
			}

			req.Calendar.Id = uint64(lastId)
		}
	}

	// TODO handle access updates

	calendar, err := s.getCalendar(ctx, userInfo, tCalendar.ID.EQ(jet.Uint64(req.Calendar.Id)))
	if err != nil {
		return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
	}

	return &CreateOrUpdateCalendarResponse{
		Calendar: calendar,
	}, nil
}

func (s *Server) DeleteCalendar(ctx context.Context, req *DeleteCalendarRequest) (*DeleteCalendarResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	check, err := s.checkIfUserHasAccessToCalendar(ctx, req.CalendarId, userInfo, calendar.AccessLevel_ACCESS_LEVEL_MANAGE)
	if err != nil {
		return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
	}
	if !check {
		return nil, errswrap.NewError(err, errorscalendar.ErrNoPerms)
	}

	stmt := tCalendar.
		UPDATE(
			tCalendar.DeletedAt,
		).
		SET(
			tCalendar.DeletedAt.SET(jet.CURRENT_TIMESTAMP()),
		).
		WHERE(tCalendar.ID.EQ(jet.Uint64(req.CalendarId)))

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, errswrap.NewError(err, errorscalendar.ErrFailedQuery)
	}

	return &DeleteCalendarResponse{}, nil
}

func (s *Server) getCalendar(ctx context.Context, userInfo *userinfo.UserInfo, condition jet.BoolExpression) (*calendar.Calendar, error) {
	stmt := tCalendar.
		SELECT(
			tCalendar.ID,
			tCalendar.CreatedAt,
			tCalendar.UpdatedAt,
			tCalendar.DeletedAt,
			tCalendar.Job,
			tCalendar.Name,
			tCalendar.Description,
			tCalendar.Public,
			tCalendar.Closed,
			tCalendar.CreatorID,
			tCalendar.CreatorJob,
			tCreator.ID,
			tCreator.Identifier,
			tCreator.Job,
			tCreator.JobGrade,
			tCreator.Firstname,
			tCreator.Lastname,
			tCreator.Dateofbirth,
			tCreator.PhoneNumber,
		).
		FROM(tCalendar.
			LEFT_JOIN(tCreator,
				tCalendar.CreatorID.EQ(tCreator.ID),
			),
		).
		GROUP_BY(tCalendar.ID)

	if condition != nil {
		stmt = stmt.WHERE(condition)

	}

	dest := &calendar.Calendar{}
	if err := stmt.QueryContext(ctx, s.db, dest); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, err
		}
	}

	if dest.Id == 0 {
		return nil, nil
	}

	if dest.Creator != nil {
		s.enricher.EnrichJobInfoSafe(userInfo, dest.Creator)
	}

	return dest, nil
}