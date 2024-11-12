package cases

import (
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type (
	Create interface {
		Run(*user.User) error
	}
	Read interface {
		Run(*user.ID) (*user.User, error)
	}
	Update interface {
		Run(*user.User, *user.Password) error
	}
	Delete interface {
		Run(*user.ID, *user.Password) error
	}
	Verify interface {
		Run(*user.ID) error
	}
	Login interface {
		Run(*user.Email, *user.Password) (*user.User, error)
	}
)
