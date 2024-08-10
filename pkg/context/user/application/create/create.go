package create

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type Create struct {
	repository.Repository
}

func (create *Create) Run(user *user.User) error {
	err := create.Repository.Save(user)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
