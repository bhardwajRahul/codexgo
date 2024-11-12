package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func Random() *User {
	id := IDWithValidValue()
	email := EmailWithValidValue()
	username := UsernameWithValidValue()
	password := PasswordWithValidValue()

	user, err := New(&Primitive{
		ID:       id.Value,
		Email:    email.Value,
		Username: username.Value,
		Password: password.Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "Random")
	}

	return user
}

func RandomPrimitive() *Primitive {
	return Random().ToPrimitive()
}
