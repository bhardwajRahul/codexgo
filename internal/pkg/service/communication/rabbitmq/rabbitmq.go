package rabbitmq

import (
	"github.com/bastean/codexgo/internal/pkg/service/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/loggers"
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communications/rabbitmq"
)

type RabbitMQ = rabbitmq.RabbitMQ

func Open(uri string, logger loggers.Logger, exchange *messages.Router, queues []*messages.Queue, consumers []messages.Consumer) (*RabbitMQ, error) {
	session, err := rabbitmq.Open(uri, logger)

	if err != nil {
		return nil, errors.BubbleUp(err, "Open")
	}

	if err = session.AddRouter(exchange); err != nil {
		return nil, errors.BubbleUp(err, "Open")
	}

	for _, queue := range queues {
		if err = session.AddQueue(queue); err != nil {
			return nil, errors.BubbleUp(err, "Open")
		}

		if err = session.AddQueueMessageBind(queue, queue.Bindings); err != nil {
			return nil, errors.BubbleUp(err, "Open")
		}
	}

	for _, consumer := range consumers {
		if err = session.AddQueueConsumer(consumer); err != nil {
			return nil, errors.BubbleUp(err, "Open")
		}
	}

	return session, nil
}

func Close(session *RabbitMQ) error {
	if err := rabbitmq.Close(session); err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
