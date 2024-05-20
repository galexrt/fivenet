package discord

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/timestamp"
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/users"
	"github.com/fivenet-app/fivenet/pkg/discord/embeds"
	"github.com/fivenet-app/fivenet/pkg/discord/modules"
	"github.com/fivenet-app/fivenet/pkg/discord/types"
	"github.com/fivenet-app/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const discordLogsEmbedChunkSize = 5

type Guild struct {
	mutex  sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc

	job string
	id  string

	logger  *zap.Logger
	bot     *Bot
	guild   *discordgo.Guild
	modules []string
}

func NewGuild(ctx context.Context, b *Bot, guild *discordgo.Guild, job string) (*Guild, error) {
	ctx, cancel := context.WithCancel(ctx)

	modules := []string{}
	if b.cfg.GroupSync.Enabled {
		modules = append(modules, "groupsync")
	}
	if b.cfg.UserInfoSync.Enabled {
		modules = append(modules, "userinfo")
	}

	return &Guild{
		mutex:  sync.Mutex{},
		ctx:    ctx,
		cancel: cancel,

		job: job,
		id:  guild.ID,

		logger:  b.logger.Named("guild").With(zap.String("job", job), zap.String("discord_guild_id", guild.ID)),
		bot:     b,
		guild:   guild,
		modules: modules,
	}, nil
}

func (g *Guild) setup() error {
	g.logger.Info("setting up guild")

	if _, err := g.bot.discord.Guild(g.guild.ID); err != nil {
		return fmt.Errorf("failed to retrieve guild info from discord api. %w", err)
	}

	return nil
}

func (g *Guild) Run() error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	start := time.Now()

	if g.guild.Unavailable {
		g.logger.Warn("discord guild is unavailable, skipping sync run")
		return nil
	}

	settings, planDiff, err := g.getSyncSettings(g.ctx, g.job)
	if err != nil {
		return err
	}
	if planDiff == nil {
		planDiff = &users.DiscordSyncChanges{}
	}

	errs := multierr.Combine()
	if settings.IsStatusLogEnabled() {
		if err := g.sendStartStatusLog(settings.StatusLogSettings.ChannelId); err != nil {
			errs = multierr.Append(errs, err)
		}
	}

	base := modules.NewBaseModule(g.logger.Named("module").With(zap.String("job", g.job), zap.String("discord_guild_id", g.guild.ID)),
		g.bot.db, g.bot.discord, g.guild, g.job, g.bot.cfg, g.bot.appCfg, g.bot.enricher, settings)

	state := &types.State{
		GuildID: g.guild.ID,
		Users:   types.Users{},
	}
	logs := []*discordgo.MessageEmbed{}
	for _, module := range g.modules {
		g.logger.Debug("running discord guild module", zap.String("dc_module", module))

		m, err := modules.GetModule(module, base)
		if err != nil {
			errs = multierr.Append(errs, fmt.Errorf("%s: %w", module, err))
			continue
		}

		s, mLogs, err := m.Plan(g.ctx)
		if err != nil {
			errs = multierr.Append(errs, fmt.Errorf("%s: %w", module, err))
			continue
		}

		state.Merge(s)
		logs = append(logs, mLogs...)
	}

	plan, ls, err := state.Calculate(g.ctx, g.bot.discord)
	if err != nil {
		errs = multierr.Append(errs, err)
		return errs
	}
	logs = append(logs, ls...)

	// Encode plan as yaml for our "change list"
	b := bytes.Buffer{}
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)
	if err := yamlEncoder.Encode(plan); err != nil {
		errs = multierr.Append(errs, err)
	}
	planDiff.Add(&users.DiscordSyncChange{
		Time: timestamp.Now(),
		Plan: b.String(),
	})

	if !settings.DryRun {
		plan.DryRun = false
		pLogs, err := plan.Apply(g.ctx, g.bot.discord)
		logs = append(logs, pLogs...)
		if err != nil {
			errs = multierr.Append(errs, err)
		}
	}

	if settings.IsStatusLogEnabled() {
		if err := g.sendStatusLog(settings.StatusLogSettings.ChannelId, logs); err != nil {
			errs = multierr.Append(errs, err)
		}

		if err := g.sendEndStatusLog(settings.StatusLogSettings.ChannelId, time.Since(start), errs); err != nil {
			errs = multierr.Append(errs, err)
		}
	}

	if err := g.setLastSyncInterval(g.ctx, g.job, planDiff); err != nil {
		g.logger.Error("error setting job props last sync time", zap.Error(err))
		errs = multierr.Append(errs, err)
	}

	g.logger.Info("completed sync run")

	return errs
}

func (g *Guild) Stop() {
	g.cancel()
}

func (g *Guild) sendStartStatusLog(channelId string) error {
	channel, err := g.bot.discord.Channel(channelId)
	if err != nil {
		return err
	}
	_ = channel

	/*
		if _, err := g.bot.discord.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed{
			Type:   discordgo.EmbedTypeRich,
			Title:  "Starting sync...",
			Author: embeds.EmbedAuthor,
			Color:  embeds.ColorInfo,
		}); err != nil {
			return err
		}*/

	return nil
}

func (g *Guild) sendStatusLog(channelId string, logs []*discordgo.MessageEmbed) error {
	if len(logs) == 0 {
		return nil
	}

	channel, err := g.bot.discord.Channel(channelId)
	if err != nil {
		return err
	}
	_ = channel

	/*
		// Split logs embeds into chunks
		for i := 0; i < len(logs); i += discordLogsEmbedChunkSize {
			end := i + discordLogsEmbedChunkSize

			if end > len(logs) {
				end = len(logs)
			}

			if _, err := g.bot.discord.ChannelMessageSendEmbeds(channel.ID, logs[i:end]); err != nil {
				return err
			}
		}*/

	return nil
}

func (g *Guild) sendEndStatusLog(channelId string, duration time.Duration, errs error) error {
	channel, err := g.bot.discord.Channel(channelId)
	if err != nil {
		return err
	}

	logs := []*discordgo.MessageEmbed{}
	if errs != nil {
		logs = append(logs, &discordgo.MessageEmbed{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Errors during sync",
			Description: fmt.Sprintf("Following errors occured during sync:\n```\n%v\n```", errs),
			Author:      embeds.EmbedAuthor,
			Color:       embeds.ColorError,
		})
	}

	logs = append(logs, &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       "Sync completed!",
		Description: fmt.Sprintf("Completed in %s.", duration),
		Author:      embeds.EmbedAuthor,
		Color:       embeds.ColorSuccess,
	})

	if _, err := g.bot.discord.ChannelMessageSendEmbeds(channel.ID, logs); err != nil {
		return err
	}

	return nil
}

func (g *Guild) getSyncSettings(ctx context.Context, job string) (*users.DiscordSyncSettings, *users.DiscordSyncChanges, error) {
	stmt := tJobProps.
		SELECT(
			tJobProps.DiscordSyncSettings,
			tJobProps.DiscordSyncChanges,
		).
		FROM(tJobProps).
		WHERE(
			tJobProps.Job.EQ(jet.String(job)),
		).
		LIMIT(1)

	var dest users.JobProps
	if err := stmt.QueryContext(ctx, g.bot.db, &dest); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, nil, err
		}
	}

	// Make sure the defaults are set
	dest.Default(job)

	return dest.DiscordSyncSettings, dest.DiscordSyncChanges, nil
}

func (g *Guild) setLastSyncInterval(ctx context.Context, job string, pDiff *users.DiscordSyncChanges) error {
	tJobProps := table.FivenetJobProps
	stmt := tJobProps.
		UPDATE(
			tJobProps.DiscordLastSync,
			tJobProps.DiscordSyncChanges,
		).
		SET(
			jet.CURRENT_TIMESTAMP(),
			pDiff,
		).
		WHERE(
			tJobProps.Job.EQ(jet.String(job)),
		)

	if _, err := stmt.ExecContext(ctx, g.bot.db); err != nil {
		return err
	}

	return nil
}
