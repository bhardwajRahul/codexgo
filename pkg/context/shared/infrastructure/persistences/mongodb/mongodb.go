package mongodb

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type MongoDB struct {
	*mongo.Client
	*mongo.Database
}

func New(uri, name string) (*MongoDB, error) {
	options := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), options)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "New",
			What:  "Failure to create a MongoDB client",
			Who:   err,
		})
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, errors.NewInternal(&errors.Bubble{
			Where: "New",
			What:  "Failure connecting to MongoDB",
			Who:   err,
		})
	}

	return &MongoDB{
		Client:   client,
		Database: client.Database(name),
	}, nil
}

func Close(ctx context.Context, mongoDB *MongoDB) error {
	if err := mongoDB.Client.Disconnect(ctx); err != nil {
		return errors.NewInternal(&errors.Bubble{
			Where: "Close",
			What:  "Failure to close connection with MongoDB",
			Who:   err,
		})
	}

	return nil
}

func HandleDuplicateKeyError(err error) error {
	re := regexp.MustCompile(`{ [A-Za-z0-9]+:`)

	rawField := re.FindString(err.Error())

	toTitle := cases.Title(language.English)

	field := toTitle.String(strings.TrimSuffix(strings.Split(rawField, " ")[1], ":"))

	return errors.NewAlreadyExist(&errors.Bubble{
		Where: "HandleDuplicateKeyError",
		What:  fmt.Sprintf("%s already registered", field),
		Why: errors.Meta{
			"Field": field,
		},
		Who: err,
	})
}

func HandleDocumentNotFound(index string, err error) error {
	return errors.NewNotExist(&errors.Bubble{
		Where: "HandleDocumentNotFound",
		What:  fmt.Sprintf("%s not found", index),
		Why: errors.Meta{
			"Index": index,
		},
		Who: err,
	})
}
