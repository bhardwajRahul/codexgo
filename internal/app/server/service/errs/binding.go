package errs

import (
	"encoding/json"
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func BindingJSON(who error) error {
	where, _, _, _ := caller.Received(caller.SkipCurrent)

	var err *json.UnmarshalTypeError

	if errors.As(who, &err) {
		return errors.New[errors.Failure](&errors.Bubble{
			Where: where,
			What:  fmt.Sprintf("Invalid type field [%s] required type is [%s] and [%s] was obtained", err.Field, err.Type, err.Value),
			Who:   who,
		})
	}

	return errors.New[errors.Internal](&errors.Bubble{
		Where: where,
		What:  "Cannot bind a JSON to a struct",
		Who:   who,
	})
}
