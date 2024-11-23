package login

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

func QueryRandomAttributes() *QueryAttributes {
	return &QueryAttributes{
		Email:    user.EmailWithValidValue().Value,
		Password: user.PlainPasswordWithValidValue().Value,
	}
}
