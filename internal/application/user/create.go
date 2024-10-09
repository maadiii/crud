package user

import (
	"context"
	"math/rand"

	"github.com/maadiii/crud/pkg/entity"
	"github.com/maadiii/crud/pkg/usecase/application/user"
	"github.com/maadiii/crud/utility/password"

	. "github.com/maadiii/regorm"
)

func Create() user.Create {
	// generate with snowflake or some thing else
	id := rand.Int63n(100000)
	return &create{User: &entity.User{Id: id}}
}

type create struct {
	*entity.User
}

func (c *create) Execute(ctx context.Context) (err error) {
	password, err := password.Hash(c.Password)
	if err != nil {
		return
	}
	c.Password = password

	return SQL(c.User).Insert(ctx, nil)
}

func (c *create) WithName(value string) user.Create {
	c.Name = value

	return c
}

func (c *create) WithFamily(value string) user.Create {
	c.Family = value

	return c
}

func (c *create) WithEmail(value string) user.Create {
	c.Email = value

	return c
}

func (c *create) WithPassword(value string) user.Create {
	c.Password = value

	return c
}

func (c *create) Id() int64 {
	return c.User.Id
}
