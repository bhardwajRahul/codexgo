package components_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
	"github.com/stretchr/testify/suite"
)

type EntityTestSuite struct {
	suite.Suite
}

func (s *EntityTestSuite) TestWithInvalidLength() {
	value, err := components.EntityWithInvalidLength()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEntity",
		What:  "Entity must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Entity": value,
		},
	}}

	s.EqualError(expected, actual.Error())
}

func (s *EntityTestSuite) TestWithInvalidAlpha() {
	value, err := components.EntityWithInvalidAlpha()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEntity",
		What:  "Entity must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Entity": value,
		},
	}}

	s.EqualError(expected, actual.Error())
}

func TestUnitEntitySuite(t *testing.T) {
	suite.Run(t, new(EntityTestSuite))
}
