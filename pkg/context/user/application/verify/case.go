package verify

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/token"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Case struct {
	role.Repository
	roles.Hasher
}

func (c *Case) Run(attributes *CommandAttributes) error {
	verifyToken, errVerifyToken := token.New(attributes.VerifyToken)
	id, errID := values.New[*values.ID](attributes.ID)
	plainPassword, errPlainPassword := values.New[*user.PlainPassword](attributes.Password)

	if err := errors.Join(errVerifyToken, errID, errPlainPassword); err != nil {
		return errors.BubbleUp(err)
	}

	aggregate, err := c.Repository.Search(&user.Criteria{
		ID: id,
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = c.Hasher.Compare(aggregate.Password.Value(), plainPassword.Value())

	if err != nil {
		return errors.BubbleUp(err)
	}

	if aggregate.IsVerified() {
		return nil
	}

	err = aggregate.ValidateVerifyToken(verifyToken)

	if err != nil {
		return errors.BubbleUp(err)
	}

	aggregate.Verified, err = values.Replace(aggregate.Verified, true)

	if err != nil {
		return errors.BubbleUp(err)
	}

	aggregate.VerifyToken = nil

	err = aggregate.UpdatedStamp()

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = c.Repository.Update(aggregate)

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
