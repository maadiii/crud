package user

import (
	"github.com/maadiii/crud/pkg/entity"
	"github.com/maadiii/crud/pkg/usecase"
)

type Create interface {
	usecase.Command

	WithName(string) Create
	WithFamily(string) Create
	WithEmail(string) Create
	WithPassword(string) Create

	Id() int64
}

type Edit interface {
	usecase.Command

	WithId(int64) Edit
	WithName(string) Edit
	WithFamily(string) Edit
}

type GetById interface {
	usecase.Command

	WithId(int64) GetById

	Id() int64
	Name() string
	Family() string
	Email() string
}

type List interface {
	usecase.Command

	SetPagination(from, count int) List
	SetSort(value string, desc bool) List
	SetSearch(key string, value any) List

	Users() []*entity.User
}
