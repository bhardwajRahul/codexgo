package collection_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/mongodb/collection"
)

type CollectionTestSuite struct {
	persistence.RepositorySuite
}

func (s *CollectionTestSuite) SetupSuite() {
	session, err := mongodb.Open(
		os.Getenv("CODEXGO_DATABASE_MONGODB_URI"),
		os.Getenv("CODEXGO_DATABASE_MONGODB_NAME"),
	)

	if err != nil {
		errors.Panic(err)
	}

	name := "users-test"

	s.RepositorySuite.SUT, err = collection.Open(session, name)

	if err != nil {
		errors.Panic(err)
	}
}

func TestIntegrationCollectionSuite(t *testing.T) {
	suite.Run(t, new(CollectionTestSuite))
}
