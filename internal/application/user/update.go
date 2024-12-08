package user

import (
	"context"

	"github.com/maadiii/crud/pkg/entity"
	"github.com/maadiii/crud/pkg/usecase/application/user"

	. "github.com/maadiii/regorm"
)

func Edit() user.Edit {
	return &edit{User: new(entity.User)}
}

type edit struct {
	*entity.User
}

func (e *edit) Execute(ctx context.Context) (err error) {
	tx, err = SQL(e.User).
		Select("name", "family").
		Where(
			EQ("id", e.Id),
		).
		UpdateTx(ctx, nil)

	SQL(e.User).SetTx(tx, true).Update(ctx)

	return
}

func (e *edit) WithId(value int64) user.Edit {
	e.Id = value

	return e
}

func (e *edit) WithName(value string) user.Edit {
	e.Name = value

	return e
}

func (e *edit) WithFamily(value string) user.Edit {
	e.Family = value

	return e
}
