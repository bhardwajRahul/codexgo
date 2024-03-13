package delete

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	valueObjectMother "github.com/bastean/codexgo/test/pkg/context/user/domain/valueObject"
)

func Create(id *sharedVO.Id) *delete.Command {
	return &delete.Command{Id: id.Value}
}

func Random() *delete.Command {
	return Create(valueObjectMother.RandomId())
}

func Invalid() *delete.Command {
	return Create(valueObjectMother.InvalidId())
}
