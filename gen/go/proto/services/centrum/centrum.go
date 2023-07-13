package centrum

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/galexrt/fivenet/gen/go/proto/resources/common"
	dispatch "github.com/galexrt/fivenet/gen/go/proto/resources/dispatch"
	"github.com/galexrt/fivenet/gen/go/proto/resources/rector"
	users "github.com/galexrt/fivenet/gen/go/proto/resources/users"
	"github.com/galexrt/fivenet/pkg/audit"
	"github.com/galexrt/fivenet/pkg/events"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/perms"
	"github.com/galexrt/fivenet/pkg/utils"
	"github.com/galexrt/fivenet/pkg/utils/dbutils"
	"github.com/galexrt/fivenet/pkg/utils/syncx"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/nats-io/nats.go"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var (
	ErrFailedQuery = status.Error(codes.Internal, "errors.CentrumService.ErrFailedQuery")
)

var (
	tCentrumUsers = table.FivenetCentrumUsers
)

type Server struct {
	CentrumServiceServer

	ctx    context.Context
	logger *zap.Logger
	tracer trace.Tracer
	db     *sql.DB
	p      perms.Permissions
	a      audit.IAuditer

	events *events.Eventus

	units      syncx.Map[string, *syncx.Map[uint64, *dispatch.Unit]]
	unitsUsers syncx.Map[int32, uint64]

	visibleJobs []string
}

func NewServer(ctx context.Context, logger *zap.Logger, tp *tracesdk.TracerProvider, db *sql.DB, p perms.Permissions, aud audit.IAuditer, e *events.Eventus, visibleJobs []string) *Server {
	return &Server{
		ctx:    ctx,
		logger: logger,

		tracer: tp.Tracer("centrum-cache"),

		db:     db,
		p:      p,
		a:      aud,
		events: e,

		units:      syncx.Map[string, *syncx.Map[uint64, *dispatch.Unit]]{},
		unitsUsers: syncx.Map[int32, uint64]{},

		visibleJobs: visibleJobs,
	}
}

func (s *Server) Start() {
	if err := s.registerEvents(); err != nil {
		s.logger.Error("failed to register events", zap.Error(err))
	}

	go func() {
		for {
			if err := s.refresh(); err != nil {
				s.logger.Error("failed to refresh centrum data", zap.Error(err))
			}

			select {
			case <-s.ctx.Done():
				return
			case <-time.After(10 * time.Second):
			}
		}
	}()
}

func (s *Server) refresh() error {
	ctx, span := s.tracer.Start(s.ctx, "centrum-refresh-cache")
	defer span.End()

	if err := s.loadUnits(ctx, 0); err != nil {
		s.logger.Error("failed to load units", zap.Error(err))
	}

	return nil
}

func (s *Server) GetSettings(ctx context.Context, req *common.EmptyRequest) (*dispatch.Settings, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	return s.getSettings(ctx, userInfo.Job)
}

func (s *Server) UpdateSettings(ctx context.Context, req *dispatch.Settings) (*dispatch.Settings, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: CentrumService_ServiceDesc.ServiceName,
		Method:  "UpdateSettings",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EVENT_TYPE_ERRORED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	if err := s.updateSettings(ctx, userInfo, req); err != nil {
		return nil, err
	}

	auditEntry.State = int16(rector.EVENT_TYPE_UPDATED)

	settings, err := s.getSettings(ctx, userInfo.Job)
	if err != nil {
		return nil, err
	}

	data, err := proto.Marshal(settings)
	if err != nil {
		return nil, err
	}
	s.broadcastToAllUnits(TopicGeneral, TypeGeneralSettings, userInfo, data)

	return settings, nil
}

func (s *Server) TakeControl(ctx context.Context, req *TakeControlRequest) (*TakeControlResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	if req.Signon {
		stmt := tCentrumUsers.
			INSERT(
				tCentrumUsers.Job,
				tCentrumUsers.UserID,
				tCentrumUsers.Identifier,
			).
			QUERY(
				tUser.
					SELECT(
						jet.String(userInfo.Job),
						jet.Int32(userInfo.UserId),
						tUser.Identifier,
					).
					FROM(tUser).
					WHERE(
						tUser.ID.EQ(jet.Int32(userInfo.UserId)),
					),
			)

		if _, err := stmt.ExecContext(ctx, s.db); err != nil {
			if !dbutils.IsDuplicateError(err) {
				return nil, err
			}
		}
	} else {
		stmt := tCentrumUsers.
			DELETE().
			WHERE(jet.AND(
				tCentrumUsers.Job.EQ(jet.String(userInfo.Job)),
				tCentrumUsers.UserID.EQ(jet.Int32(userInfo.UserId)),
			)).
			LIMIT(1)

		if _, err := stmt.ExecContext(ctx, s.db); err != nil {
			return nil, err
		}
	}

	controllers, err := s.getControllers(ctx, userInfo.Job)
	if err != nil {
		return nil, err
	}

	settings, err := s.getSettings(ctx, userInfo.Job)
	if err != nil {
		return nil, err
	}
	// If center is enabled update settings active state accordingly
	if settings.Enabled {
		settings.Active = len(controllers) > 0

		if err := s.updateSettings(ctx, userInfo, settings); err != nil {
			return nil, err
		}
	}

	data, err := proto.Marshal(&ControllerChange{
		Controllers: controllers,
		Active:      settings.Active,
	})
	if err != nil {
		return nil, err
	}
	s.broadcastToAllUnits(TopicGeneral, TypeGeneralControllers, userInfo, data)

	return nil, nil
}

func (s *Server) Stream(req *StreamRequest, srv CentrumService_StreamServer) error {
	userInfo := auth.MustGetUserInfoFromContext(srv.Context())

	unitId, err := s.getUnitIDFromUserID(srv.Context(), userInfo.UserId)
	if err != nil {
		return err
	}
	settings, err := s.getSettings(srv.Context(), userInfo.Job)
	if err != nil {
		return err
	}

	msgCh := make(chan *nats.Msg, 128)

	controllers, err := s.getControllers(srv.Context(), userInfo.Job)
	if err != nil {
		return err
	}

	if utils.InSliceFunc(controllers, func(in *users.UserShort) bool {
		return in.UserId == userInfo.UserId
	}) {
		sub, err := s.events.JS.ChanSubscribe(fmt.Sprintf(BaseSubject+".%s.*.*.%d", userInfo.Job, unitId), msgCh)
		if err != nil {
			return err
		}
		defer sub.Unsubscribe()
	} else {
		sub, err := s.events.JS.ChanSubscribe(fmt.Sprintf(BaseSubject+".%s.>", userInfo.Job), msgCh)
		if err != nil {
			return err
		}
		defer sub.Unsubscribe()
	}

	unit, _ := s.getUnit(srv.Context(), userInfo, unitId)

	// Send initial message to client
	resp := &StreamResponse{}
	resp.Change = &StreamResponse_Initial{
		Initial: &Initial{
			Controller: true,
			Settings:   settings,
			Unit:       unit,
		},
	}

	if err := srv.Send(resp); err != nil {
		return err
	}

	for {
		select {
		case <-srv.Context().Done():
			return nil
		case msg := <-msgCh:
			msg.Ack()
			topic, tType := s.getEventTypeFromSubject(msg.Subject)

			switch topic {
			case TopicGeneral:
				switch tType {
				case TypeGeneralControllers:
					var dest ControllerChange
					if err := proto.Unmarshal(msg.Data, &dest); err != nil {
						return err
					}

					resp.Change = &StreamResponse_Controllers{
						Controllers: &dest,
					}

				case TypeGeneralSettings:
					var dest dispatch.Settings
					if err := proto.Unmarshal(msg.Data, &dest); err != nil {
						return err
					}

					resp.Change = &StreamResponse_Settings{
						Settings: &dest,
					}
				}

			case TopicDispatch:
				switch tType {
				case TypeDispatchUpdated:
					var dest dispatch.Dispatch
					if err := proto.Unmarshal(msg.Data, &dest); err != nil {
						return err
					}

					resp.Change = &StreamResponse_DispatchUpdate{
						DispatchUpdate: &dest,
					}

				case TypeDispatchStatus:
					var dest dispatch.DispatchStatus
					if err := proto.Unmarshal(msg.Data, &dest); err != nil {
						return err
					}

					resp.Change = &StreamResponse_DispatchStatus{
						DispatchStatus: &dest,
					}

				case TypeDispatchAssigned:
					var dest dispatch.Dispatch
					if err := proto.Unmarshal(msg.Data, &dest); err != nil {
						return err
					}

					resp.Change = &StreamResponse_DispatchAssigned{
						DispatchAssigned: &dest,
					}

				case TypeDispatchUnassigned:
					var dest dispatch.Dispatch
					if err := proto.Unmarshal(msg.Data, &dest); err != nil {
						return err
					}

					resp.Change = &StreamResponse_DispatchUnassigned{
						DispatchUnassigned: &dest,
					}
				}

			case TopicUnit:
				switch tType {
				case TypeUnitUpdated:
					var dest dispatch.Dispatch
					if err := proto.Unmarshal(msg.Data, &dest); err != nil {
						return err
					}

					resp.Change = &StreamResponse_DispatchAssigned{
						DispatchAssigned: &dest,
					}

				case TypeUnitStatus:
					var dest dispatch.UnitStatus
					if err := proto.Unmarshal(msg.Data, &dest); err != nil {
						return err
					}

					resp.Change = &StreamResponse_UnitStatus{
						UnitStatus: &dest,
					}
				case TypeUnitAssigned:
					var dest dispatch.Unit
					if err := proto.Unmarshal(msg.Data, &dest); err != nil {
						return err
					}

					found := false
					for _, u := range dest.Users {
						if u.UserId == userInfo.UserId {
							found = true
							break
						}
					}

					if found {
						resp.Change = &StreamResponse_UnitAssigned{
							UnitAssigned: &dest,
						}
					} else {
						resp.Change = &StreamResponse_UnitAssigned{
							UnitAssigned: nil,
						}
					}

				case TypeUnitDeleted:
					var dest dispatch.Unit
					if err := proto.Unmarshal(msg.Data, &dest); err != nil {
						return err
					}

					resp.Change = &StreamResponse_UnitDeleted{
						UnitDeleted: dest.Id,
					}
				}
			}
		}

		if err := srv.Send(resp); err != nil {
			return err
		}
	}
}
