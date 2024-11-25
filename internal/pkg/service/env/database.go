package env

import (
	"os"
)

var (
	DatabaseMongoDBURI, DatabaseMongoDBName string
)

var (
	DatabaseSQLiteDSN string
)

func Database() {
	DatabaseMongoDBURI = os.Getenv(DATABASE_MONGODB_URI)
	DatabaseMongoDBName = os.Getenv(DATABASE_MONGODB_NAME)

	DatabaseSQLiteDSN = os.Getenv(DATABASE_SQLITE_DSN)
}

func HasDatabase() bool {
	return DatabaseMongoDBURI != ""
}
