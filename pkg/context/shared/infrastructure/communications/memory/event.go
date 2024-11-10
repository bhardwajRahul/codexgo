package memory

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

type (
	eventMapper = map[events.Key][]events.Consumer
)

type EventBus struct {
	Consumers eventMapper
}

func (bus *EventBus) Subscribe(key events.Key, consumer events.Consumer) error {
	bus.Consumers[key] = append(bus.Consumers[key], consumer)
	return nil
}

func (bus *EventBus) Publish(event *events.Event) error {
	consumers, ok := bus.Consumers[event.Key]

	if !ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Publish",
			What:  "Failure to execute a Event without a Consumer",
			Why: errors.Meta{
				"Event": event.Key,
			},
		})
	}

	for _, consumer := range consumers {
		err := consumer.On(event)

		if err != nil {
			return errors.BubbleUp(err, "Publish")
		}
	}

	return nil
}

func NewEventBus(mapper eventMapper) (*EventBus, error) {
	bus := &EventBus{
		Consumers: make(eventMapper, len(mapper)),
	}

	var err error

	for key, consumers := range mapper {
		for _, consumer := range consumers {
			err = bus.Subscribe(key, consumer)

			if err != nil {
				return nil, errors.BubbleUp(err, "NewEventBus")
			}
		}
	}

	return bus, nil
}
