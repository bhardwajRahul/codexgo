package event

import (
	"reflect"

	"github.com/bastean/codexgo/v4/pkg/context/notification/application/confirmation"
	"github.com/bastean/codexgo/v4/pkg/context/notification/application/password"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

var UserCreatedSucceededRecipient, _ = values.New[*messages.Recipient](messages.FormatRecipient(&messages.RecipientComponents{
	Service: "user",
	Entity:  "user",
	Trigger: "send_confirmation",
	Action:  "created",
	Status:  messages.Status.Succeeded,
}))

var UserResetQueuedRecipient, _ = values.New[*messages.Recipient](messages.FormatRecipient(&messages.RecipientComponents{
	Service: "user",
	Entity:  "user",
	Trigger: "send_reset",
	Action:  "reset",
	Status:  messages.Status.Queued,
}))

var RabbitMQueueMapper = rabbitmq.Mapper{
	user.CreatedSucceededKey.Value(): []*rabbitmq.Queue{
		{
			Name:       UserCreatedSucceededRecipient,
			BindingKey: user.CreatedSucceededKey.Value(),
			Attributes: reflect.TypeOf(new(confirmation.EventAttributes)),
			Meta:       reflect.TypeOf(new(confirmation.EventMeta)),
		},
	},
	user.ResetQueuedKey.Value(): []*rabbitmq.Queue{
		{
			Name:       UserResetQueuedRecipient,
			BindingKey: user.ResetQueuedKey.Value(),
			Attributes: reflect.TypeOf(new(password.EventAttributes)),
			Meta:       reflect.TypeOf(new(password.EventMeta)),
		},
	},
}
