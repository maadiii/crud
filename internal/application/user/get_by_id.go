package user

import (
	"context"

	"github.com/maadiii/crud/pkg/entity"
	"github.com/maadiii/crud/pkg/usecase/application/user"

	. "github.com/maadiii/regorm"
)

func GetById() user.GetById {
	return &getById{User: new(entity.User)}
}

type getById struct {
	*entity.User
}

func (g *getById) Execute(ctx context.Context) (err error) {
	g.User, err = SQL(g.User).
		Select("id", "name", "family", "email").
		Where(
			EQ("id", g.User.Id),
		).
		One(ctx)

	return
}

func (g *getById) WithId(value int64) user.GetById {
	g.User.Id = value

	return g
}

func (g *getById) Id() int64 {
	return g.User.Id
}

func (g *getById) Name() string {
	return g.User.Name
}

func (g *getById) Family() string {
	return g.User.Family
}

func (g *getById) Email() string {
	return g.User.Email
}
