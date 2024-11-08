package update_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type UpdateTestSuite struct {
	suite.Suite
	sut        commands.Handler
	update     cases.Update
	hashing    *cryptographic.HashingMock
	repository *persistence.UserMock
}

func (suite *UpdateTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.update = &update.Update{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &update.Handler{
		Update: suite.update,
	}
}

func (suite *UpdateTestSuite) TestSubscribedTo() {
	const expected commands.Type = "user.command.updating.user"

	actual := suite.sut.SubscribedTo()

	suite.Equal(expected, actual)
}

func (suite *UpdateTestSuite) TestHandle() {
	command := update.RandomCommand()

	account, err := user.FromPrimitive(&user.Primitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.UpdatedPassword,
	})

	suite.NoError(err)

	id, err := user.NewId(command.Id)

	suite.NoError(err)

	criteria := &repository.SearchCriteria{
		Id: id,
	}

	suite.repository.On("Search", criteria).Return(account)

	suite.hashing.On("IsNotEqual", account.Password.Value, command.Password).Return(false)

	suite.repository.On("Update", account)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())
}

func TestUnitUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
