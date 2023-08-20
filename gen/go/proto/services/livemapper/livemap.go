package livemapper

import (
	"context"
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/galexrt/fivenet/gen/go/proto/resources/livemap"
	"github.com/galexrt/fivenet/gen/go/proto/resources/timestamp"
	users "github.com/galexrt/fivenet/gen/go/proto/resources/users"
	"github.com/galexrt/fivenet/pkg/config"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/grpc/auth/userinfo"
	"github.com/galexrt/fivenet/pkg/mstlystcdata"
	"github.com/galexrt/fivenet/pkg/perms"
	"github.com/galexrt/fivenet/pkg/tracker"
	"github.com/galexrt/fivenet/pkg/utils"
	"github.com/galexrt/fivenet/pkg/utils/syncx"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	DispatchMarkerLimit = 60
)

var (
	ErrStreamFailed = status.Error(codes.Internal, "errors.LivemapperService.ErrStreamFailed")
)

var (
	tLocs     = table.FivenetUserLocations
	tUsers    = table.Users.AS("user")
	tJobProps = table.FivenetJobProps
)

type Server struct {
	LivemapperServiceServer

	ctx      context.Context
	logger   *zap.Logger
	tracer   trace.Tracer
	db       *sql.DB
	ps       perms.Permissions
	enricher *mstlystcdata.Enricher
	tracker  *tracker.Tracker

	dispatchesCache syncx.Map[string, []*livemap.DispatchMarker]
	usersCache      syncx.Map[string, []*livemap.UserMarker]

	broker *utils.Broker[interface{}]

	refreshTime time.Duration
	visibleJobs []string
}

type Params struct {
	fx.In

	LC fx.Lifecycle

	Logger   *zap.Logger
	TP       *tracesdk.TracerProvider
	DB       *sql.DB
	Perms    perms.Permissions
	Enricher *mstlystcdata.Enricher
	Config   *config.Config
	Tracker  *tracker.Tracker
}

func NewServer(p Params) *Server {
	ctx, cancel := context.WithCancel(context.Background())

	broker := utils.NewBroker[interface{}](ctx)
	s := &Server{
		ctx:    ctx,
		logger: p.Logger,

		tracer:   p.TP.Tracer("livemapper-cache"),
		db:       p.DB,
		ps:       p.Perms,
		enricher: p.Enricher,
		tracker:  p.Tracker,

		dispatchesCache: syncx.Map[string, []*livemap.DispatchMarker]{},
		usersCache:      syncx.Map[string, []*livemap.UserMarker]{},

		broker: broker,

		refreshTime: p.Config.Game.Livemap.RefreshTime,
		visibleJobs: p.Config.Game.Livemap.Jobs,
	}

	p.LC.Append(fx.StartHook(func(_ context.Context) error {
		go broker.Start()
		go s.start()
		return nil
	}))
	p.LC.Append(fx.StopHook(func(_ context.Context) error {
		cancel()
		return nil
	}))

	return s
}

func (s *Server) start() {
	for {
		s.refreshCache()

		select {
		case <-s.ctx.Done():
			s.broker.Stop()
			return
		case <-time.After(s.refreshTime):
		}
	}
}

func (s *Server) refreshCache() {
	ctx, span := s.tracer.Start(s.ctx, "livemap-refresh-cache")
	defer span.End()

	if err := s.refreshUserLocations(ctx); err != nil {
		s.logger.Error("failed to refresh livemap users cache", zap.Error(err))
	}
	//if err := s.refreshDispatches(ctx); err != nil {
	//	s.logger.Error("failed to refresh livemap dispatches cache", zap.Error(err))
	//}

	s.broker.Publish(nil)
}

func (s *Server) Stream(req *StreamRequest, srv LivemapperService_StreamServer) error {
	userInfo := auth.MustGetUserInfoFromContext(srv.Context())

	playersAttr, err := s.ps.Attr(userInfo, LivemapperServicePerm, LivemapperServiceStreamPerm, LivemapperServiceStreamPlayersPermField)
	if err != nil {
		return ErrStreamFailed
	}

	var playersJobs map[string]int32
	if playersAttr != nil {
		playersJobs, _ = playersAttr.(map[string]int32)
	}

	if userInfo.SuperUser {
		playersJobs = map[string]int32{}
		for _, j := range s.visibleJobs {
			playersJobs[j] = -1
		}
	}

	resp := &StreamResponse{}

	if len(playersJobs) == 0 {
		if err := srv.Send(resp); err != nil {
			return err
		}

		return nil
	}

	resp.JobsUsers = []*users.Job{}
	for job := range playersJobs {
		j := &users.Job{
			Name: job,
		}
		s.enricher.EnrichJobName(j)
		resp.JobsUsers = append(resp.JobsUsers, j)
	}

	signalCh := s.broker.Subscribe()
	defer s.broker.Unsubscribe(signalCh)

	for {
		userMarkers, _, err := s.getUserLocations(playersJobs, userInfo)
		if err != nil {
			return ErrStreamFailed
		}
		resp.Users = userMarkers

		if err := srv.Send(resp); err != nil {
			return err
		}

		select {
		case <-srv.Context().Done():
			return nil
		case <-signalCh:
		}
	}
}

func (s *Server) getUserLocations(jobs map[string]int32, userInfo *userinfo.UserInfo) ([]*livemap.UserMarker, bool, error) {
	ds := []*livemap.UserMarker{}

	found := false
	if userInfo.SuperUser {
		found = true
	}
	for job, grade := range jobs {
		markers, ok := s.tracker.GetUsers(job)
		if !ok {
			continue
		}

		markers.Range(func(key int32, marker *livemap.UserMarker) bool {
			// SuperUser returns grade as `-1`, job has access to that grade or it is the user itself
			if grade == -1 || (marker.User.JobGrade <= grade || key == userInfo.UserId) {
				ds = append(ds, marker)
			}

			// If the user is found in the list of user markers, set found state
			if !found && (userInfo.Job == job && key == userInfo.UserId) {
				found = true
			}

			return true
		})
	}

	if found {
		return ds, true, nil
	}

	return nil, false, nil
}

func (s *Server) getUserDispatches(jobs []string) ([]*livemap.DispatchMarker, error) {
	ds := []*livemap.DispatchMarker{}

	for _, job := range jobs {
		markers, ok := s.dispatchesCache.Load(job)
		if !ok {
			continue
		}

		ds = append(ds, markers...)
	}

	return ds, nil
}

func (s *Server) refreshUserLocations(ctx context.Context) error {
	markers := map[string][]*livemap.UserMarker{}

	tLocs := tLocs.AS("genericmarker")
	stmt := tLocs.
		SELECT(
			tLocs.Identifier,
			tLocs.Job,
			tLocs.X,
			tLocs.Y,
			tLocs.UpdatedAt,
			tUsers.ID.AS("genericmarker.id"),
			tUsers.ID.AS("usermarker.user_id"),
			tUsers.ID.AS("user.id"),
			tUsers.Identifier,
			tUsers.Job,
			tUsers.JobGrade,
			tUsers.Firstname,
			tUsers.Lastname,
			tJobProps.LivemapMarkerColor.AS("genericmarker.color"),
		).
		FROM(
			tLocs.
				INNER_JOIN(tUsers,
					tLocs.Identifier.EQ(tUsers.Identifier),
				).
				LEFT_JOIN(tJobProps,
					tJobProps.Job.EQ(tUsers.Job),
				),
		).
		WHERE(jet.AND(
			tLocs.Hidden.IS_FALSE(),
			tLocs.UpdatedAt.GT_EQ(jet.CURRENT_TIMESTAMP().SUB(jet.INTERVAL(60, jet.MINUTE))),
		))

	var dest []*livemap.UserMarker
	if err := stmt.QueryContext(ctx, s.db, &dest); err != nil {
		return err
	}

	for i := 0; i < len(dest); i++ {
		s.enricher.EnrichJobInfo(dest[i].User)

		job := dest[i].User.Job
		if _, ok := markers[job]; !ok {
			markers[job] = []*livemap.UserMarker{}
		}
		if dest[i].Marker.Color == nil {
			defaultColor := users.DefaultLivemapMarkerColor
			dest[i].Marker.Color = &defaultColor
		}

		markers[job] = append(markers[job], dest[i])
	}
	for job, v := range markers {
		s.usersCache.Store(job, v)
	}

	return nil
}

func (s *Server) refreshDispatches(ctx context.Context) error {
	if len(s.visibleJobs) == 0 {
		s.logger.Warn("empty livemap jobs in config, no dispatches can be found because of that")
		return nil
	}

	gksphoneJobM := table.GksphoneJobMessage
	stmt := gksphoneJobM.
		SELECT(
			gksphoneJobM.ID,
			gksphoneJobM.Name,
			gksphoneJobM.Number,
			gksphoneJobM.Message,
			gksphoneJobM.Gps,
			gksphoneJobM.Owner,
			gksphoneJobM.Jobm,
			gksphoneJobM.Anon,
			gksphoneJobM.Time,
		).
		FROM(
			gksphoneJobM,
		).
		WHERE(
			jet.AND(
				gksphoneJobM.Jobm.REGEXP_LIKE(jet.String("\\[\"("+strings.Join(s.visibleJobs, "|")+")\"\\]")),
				gksphoneJobM.Time.GT_EQ(jet.CURRENT_TIMESTAMP().SUB(jet.INTERVAL(20, jet.MINUTE))),
			),
		).
		ORDER_BY(
			gksphoneJobM.Owner.ASC(),
			gksphoneJobM.Time.DESC(),
		).
		LIMIT(DispatchMarkerLimit)

	var dest []*model.GksphoneJobMessage
	if err := stmt.QueryContext(ctx, s.db, &dest); err != nil {
		return err
	}

	markers := map[string][]*livemap.DispatchMarker{}
	for _, d := range dest {
		gps, _ := strings.CutPrefix(*d.Gps, "GPS: ")
		gpsSplit := strings.Split(gps, ", ")
		x, _ := strconv.ParseFloat(gpsSplit[0], 32)
		y, _ := strconv.ParseFloat(gpsSplit[1], 32)

		var icon string
		var color string
		if d.Owner == 0 {
			icon = "dispatch-open.svg"
			color = "96E6B3"
		} else {
			icon = "dispatch-closed.svg"
			color = "DA3E52"
		}

		var name string
		if d.Anon != nil && *d.Anon == "1" {
			name = "Anonym"
		} else {
			name = *d.Name
		}

		var message string
		if d.Message != nil && *d.Message != "" {
			message = *d.Message
		} else {
			message = "N/A"
		}

		// Remove the "json" leftovers (the data looks like this, e.g., `["ambulance"]`)
		job := strings.TrimSuffix(strings.TrimPrefix(*d.Jobm, "[\""), "\"]")
		if _, ok := markers[job]; !ok {
			markers[job] = []*livemap.DispatchMarker{}
		}
		marker := &livemap.DispatchMarker{
			Marker: &livemap.GenericMarker{
				Id:        uint64(d.ID),
				X:         x,
				Y:         y,
				Color:     &color,
				Icon:      &icon,
				Name:      name,
				Popup:     &message,
				UpdatedAt: timestamp.New(d.Time),
			},
			Job: job,
		}
		if d.Owner == 0 {
			marker.Active = true
		}

		s.enricher.EnrichJobName(marker)
		markers[job] = append(markers[job], marker)
	}

	for job, v := range markers {
		s.dispatchesCache.Store(job, v)
	}

	return nil
}
