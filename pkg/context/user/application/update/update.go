package update

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/hashes"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type Case struct {
	repository.Repository
	hashes.Hashing
}

func (c *Case) Run(aggregate *user.User, updated *user.PlainPassword) error {
	account, err := c.Repository.Search(&repository.SearchCriteria{
		ID: aggregate.ID,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = hashes.IsPasswordInvalid(c.Hashing, account.CipherPassword.Value, aggregate.PlainPassword.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	if updated != nil {
		hashed, err := c.Hashing.Hash(updated.Value)

		if err != nil {
			return errors.BubbleUp(err, "Run")
		}

		aggregate.CipherPassword, err = user.NewCipherPassword(hashed)

		if err != nil {
			return errors.BubbleUp(err, "Run")
		}
	}

	aggregate.Verified = account.Verified

	err = c.Repository.Update(aggregate)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
