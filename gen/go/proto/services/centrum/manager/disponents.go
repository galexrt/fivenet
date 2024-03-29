package manager

import (
	"context"

	"github.com/galexrt/fivenet/gen/go/proto/resources/centrum"
	errorscentrum "github.com/galexrt/fivenet/gen/go/proto/services/centrum/errors"
	eventscentrum "github.com/galexrt/fivenet/gen/go/proto/services/centrum/events"
	"github.com/galexrt/fivenet/pkg/grpc/errswrap"
	"github.com/galexrt/fivenet/pkg/utils/dbutils"
	jet "github.com/go-jet/jet/v2/mysql"
	"google.golang.org/protobuf/proto"
)

func (s *Manager) DisponentSignOn(ctx context.Context, job string, userId int32, signon bool) error {
	if signon {
		if _, ok := s.tracker.GetUserById(userId); !ok {
			return errorscentrum.ErrNotOnDuty
		}

		stmt := tCentrumUsers.
			INSERT(
				tCentrumUsers.Job,
				tCentrumUsers.UserID,
				tCentrumUsers.Identifier,
			).
			VALUES(
				job,
				userId,
				tUsers.
					SELECT(
						tUsers.Identifier.AS("identifier"),
					).
					FROM(tUsers).
					WHERE(
						tUsers.ID.EQ(jet.Int32(userId)),
					).
					LIMIT(1),
			)

		if _, err := stmt.ExecContext(ctx, s.db); err != nil {
			if !dbutils.IsDuplicateError(err) {
				return errswrap.NewError(errorscentrum.ErrFailedQuery, err)
			}
		}
	} else {
		stmt := tCentrumUsers.
			DELETE().
			WHERE(jet.AND(
				tCentrumUsers.Job.EQ(jet.String(job)),
				tCentrumUsers.UserID.EQ(jet.Int32(userId)),
			)).
			LIMIT(1)

		if _, err := stmt.ExecContext(ctx, s.db); err != nil {
			return errswrap.NewError(errorscentrum.ErrFailedQuery, err)
		}
	}

	// Load updated disponents into state
	if err := s.LoadDisponentsFromDB(ctx, job); err != nil {
		return errswrap.NewError(errorscentrum.ErrFailedQuery, err)
	}

	disponents, err := s.GetDisponents(job)
	if err != nil {
		return errorscentrum.ErrFailedQuery
	}

	change := &centrum.Disponents{
		Job:        job,
		Disponents: disponents,
	}
	data, err := proto.Marshal(change)
	if err != nil {
		return errswrap.NewError(errorscentrum.ErrFailedQuery, err)
	}

	if _, err := s.js.Publish(eventscentrum.BuildSubject(eventscentrum.TopicGeneral, eventscentrum.TypeGeneralDisponents, job), data); err != nil {
		return err
	}

	return nil
}
