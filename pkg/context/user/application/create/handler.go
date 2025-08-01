package create

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

var CommandKey, _ = values.New[*messages.Key](messages.FormatKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Command,
	Entity:  "user",
	Action:  "create",
	Status:  messages.Status.Queued,
}))

type CommandAttributes = struct {
	VerifyToken, ID, Email, Username, Password string
}

type CommandMeta = struct{}

type Handler struct {
	*Case
}

func (h *Handler) Handle(command *messages.Message) error {
	attributes, ok := command.Attributes.(*CommandAttributes)

	if !ok {
		return errors.CommandAssertion()
	}

	err := h.Case.Run(attributes)

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
