package eventscentrum

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/galexrt/fivenet/pkg/events"
	natsutils "github.com/galexrt/fivenet/pkg/nats"
	"github.com/nats-io/nats.go"
)

const (
	BaseSubject events.Subject = "centrum"

	TopicGeneral          events.Topic = "general"
	TypeGeneralSettings   events.Type  = "settings"
	TypeGeneralDisponents events.Type  = "disponents"

	TopicDispatch       events.Topic = "dispatch"
	TypeDispatchDeleted events.Type  = "deleted"
	TypeDispatchCreated events.Type  = "created"
	TypeDispatchUpdated events.Type  = "updated"
	TypeDispatchStatus  events.Type  = "status"

	TopicUnit       events.Topic = "unit"
	TypeUnitCreated events.Type  = "created"
	TypeUnitDeleted events.Type  = "deleted"
	TypeUnitUpdated events.Type  = "updated"
	TypeUnitStatus  events.Type  = "status"
)

func SplitSubject(subject string) (string, events.Topic, events.Type) {
	split := strings.Split(subject, ".")
	if len(split) < 3 {
		return "", "", ""
	}

	return split[1], events.Topic(split[2]), events.Type(split[3])
}

// Structure: "BASE_SUJBECT.JOB.TOPIC.TYPE.ID"
func BuildSubject(topic events.Topic, tType events.Type, job string) string {
	return fmt.Sprintf("%s.%s.%s.%s", BaseSubject, job, topic, tType)
}

func RegisterStream(ctx context.Context, js nats.JetStreamContext) error {
	cfg := &nats.StreamConfig{
		Name:        "CENTRUM",
		Description: natsutils.Description,
		Retention:   nats.InterestPolicy,
		Subjects:    []string{fmt.Sprintf("%s.>", BaseSubject)},
		Discard:     nats.DiscardOld,
		MaxAge:      120 * time.Second,
		Storage:     nats.MemoryStorage,
		Duplicates:  20 * time.Second,
	}
	if _, err := natsutils.CreateOrUpdateStream(ctx, js, cfg); err != nil {
		return err
	}

	return nil
}
