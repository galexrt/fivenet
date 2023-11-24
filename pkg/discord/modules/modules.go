package modules

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/galexrt/fivenet/gen/go/proto/resources/users"
	"github.com/galexrt/fivenet/pkg/config"
	"github.com/galexrt/fivenet/pkg/mstlystcdata"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"go.uber.org/zap"
)

var (
	tOauth2Accs = table.FivenetOauth2Accounts
	tAccs       = table.FivenetAccounts
	tUsers      = table.Users.AS("users")
	tJobProps   = table.FivenetJobProps.AS("jobprops")
)

var Modules = map[string]NewModuleFunc{}

type NewModuleFunc func(*BaseModule) (Module, error)

type Module interface {
	Run() error
}

func GetModule(name string, base *BaseModule) (Module, error) {
	fn, ok := Modules[name]
	if !ok {
		return nil, fmt.Errorf("no module found by name %s", name)
	}

	// "Wrap" logger with module name
	base.logger = base.logger.Named(name)

	return fn(base)
}

type BaseModule struct {
	ctx     context.Context
	logger  *zap.Logger
	db      *sql.DB
	discord *discordgo.Session
	guild   *discordgo.Guild
	job     string
	cfg     *config.DiscordBot

	enricher *mstlystcdata.Enricher
}

func NewBaseModule(ctx context.Context, logger *zap.Logger, db *sql.DB, discord *discordgo.Session, guild *discordgo.Guild, job string, cfg *config.DiscordBot, enricher *mstlystcdata.Enricher) *BaseModule {
	return &BaseModule{
		ctx:     ctx,
		logger:  logger,
		db:      db,
		discord: discord,
		guild:   guild,
		job:     job,
		cfg:     cfg,

		enricher: enricher,
	}
}

func (g *BaseModule) GetSyncSettings(ctx context.Context, job string) (*users.DiscordSyncSettings, error) {
	stmt := tJobProps.
		SELECT(
			tJobProps.DiscordSyncSettings,
		).
		FROM(tJobProps).
		WHERE(
			tJobProps.Job.EQ(jet.String(job)),
		).
		LIMIT(1)

	var dest users.JobProps
	if err := stmt.QueryContext(ctx, g.db, &dest); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, err
		}
	}

	// Make sure the defaults are set
	dest.Default(job)

	return dest.DiscordSyncSettings, nil
}