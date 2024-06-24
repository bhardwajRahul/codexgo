package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueobjs"
	"github.com/stretchr/testify/suite"
)

type ServiceValueObjectTestSuite struct {
	suite.Suite
}

func (suite *ServiceValueObjectTestSuite) SetupTest() {}

func (suite *ServiceValueObjectTestSuite) TestWithInvalidLength() {
	service, err := valueobjs.ServiceWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewService",
		What:  "service must be between " + "1" + " to " + "20" + " characters and be alphanumeric only",
		Why: errors.Meta{
			"Service": service,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *ServiceValueObjectTestSuite) TestWithInvalidAlphanumeric() {
	service, err := valueobjs.ServiceWithInvalidAlphanumeric()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewService",
		What:  "service must be between " + "1" + " to " + "20" + " characters and be alphanumeric only",
		Why: errors.Meta{
			"Service": service,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitServiceValueObjectSuite(t *testing.T) {
	suite.Run(t, new(ServiceValueObjectTestSuite))
}
