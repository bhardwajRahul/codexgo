package caller_test

import (
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type ParseTestSuite struct {
	suite.Default
}

func (s *ParseTestSuite) TestSentinel() {
	s.Equal(`[^.()[\]]+`, caller.RExAvoid)
}

func (s *ParseTestSuite) TestParseWithGenericPointerReceiver() {
	pkg, receiver, method := caller.Mother().ParseValidValues()

	received := fmt.Sprintf("%s.(*%s[...]).%s", pkg, receiver, method)

	actual := caller.Parse(received)

	expected := []string{pkg, "*" + receiver, method}

	s.Equal(expected, actual)
}

func (s *ParseTestSuite) TestParseWithPointerReceiver() {
	pkg, receiver, method := caller.Mother().ParseValidValues()

	received := fmt.Sprintf("%s.(*%s).%s", pkg, receiver, method)

	actual := caller.Parse(received)

	expected := []string{pkg, "*" + receiver, method}

	s.Equal(expected, actual)
}

func (s *ParseTestSuite) TestParseWithReceiver() {
	pkg, receiver, method := caller.Mother().ParseValidValues()

	received := fmt.Sprintf("%s.%s.%s", pkg, receiver, method)

	actual := caller.Parse(received)

	expected := []string{pkg, receiver, method}

	s.Equal(expected, actual)
}

func (s *ParseTestSuite) TestParseWithGenericFunction() {
	pkg, _, function := caller.Mother().ParseValidValues()

	received := fmt.Sprintf("%s.%s[...]", pkg, function)

	actual := caller.Parse(received)

	expected := []string{pkg, function}

	s.Equal(expected, actual)
}

func (s *ParseTestSuite) TestParseWithFunction() {
	pkg, _, function := caller.Mother().ParseValidValues()

	received := fmt.Sprintf("%s.%s", pkg, function)

	actual := caller.Parse(received)

	expected := []string{pkg, function}

	s.Equal(expected, actual)
}

func TestUnitParseSuite(t *testing.T) {
	suite.Run(t, new(ParseTestSuite))
}
