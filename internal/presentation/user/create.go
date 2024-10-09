package user

import (
	"context"

	"github.com/maadiii/crud/internal/application/user"

	"github.com/maadiii/hertz/server"
)

func init() {
	server.Register(Create)
}

// [POST] /api/v1/users 201 json
func Create(ctx context.Context, req *server.Request, in *ReqUserCreate) (out *ResUserCreate, err error) {
	command := user.Create().
		WithName(in.Name).
		WithFamily(in.Family).
		WithEmail(in.Email).
		WithPassword(in.Password)
	err = command.Execute(ctx)

	out = &ResUserCreate{Id: command.Id()}

	return
}

type ReqUserCreate struct {
	Name     string `json:"name"`
	Family   string `json:"family"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResUserCreate struct {
	Id int64 `json:"id"`
}
