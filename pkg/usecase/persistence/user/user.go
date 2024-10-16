package user

import (
	"context"

	"github.com/maadiii/crud/pkg/entity"
)

type Persistence interface {
	Insert(context.Context, *entity.User) error
	Update(context.Context, *entity.User) error
	Delete(context.Context, int64) error
	GetById(context.Context, *entity.User) error
	GetByEmail(context.Context, *entity.User) error
}
