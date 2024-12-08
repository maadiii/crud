package user

import (
	"context"

	"github.com/maadiii/crud/internal/application/user"
	usecase "github.com/maadiii/crud/pkg/usecase/application/user"

	"github.com/maadiii/hertz/server"
)

func init() {
	server.Register(Create)
}

// [POST] /api/v1/users 201 json
func Create(ctx context.Context, req *server.Request, in *ReqUserCreate) (out *ResUserCreate, err error) {
	command := in.mapRequest()
	err = command.Execute(ctx)

	out.mapResponse(command)

	return
}

type ReqUserCreate struct {
	Name     string `json:"name"`
	Family   string `json:"family"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *ReqUserCreate) mapRequest() usecase.Create {
	return user.Create().
		WithName(r.Name).
		WithFamily(r.Family).
		WithEmail(r.Email).
		WithPassword(r.Password)

}

type ResUserCreate struct {
	Id int64 `json:"id"`
}

func (r *ResUserCreate) mapResponse(command usecase.Create) *ResUserCreate {
	return &ResUserCreate{
		Id: command.Id(),
	}
}
