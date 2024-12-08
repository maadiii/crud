package user

import (
	"context"

	"github.com/maadiii/crud/internal/application/user"
	usecase "github.com/maadiii/crud/pkg/usecase/application/user"

	"github.com/maadiii/hertz/server"
)

func init() {
	server.Register(GetById)
}

// [GET] /api/v1/users/:id 200 json
func GetById(ctx context.Context, req *server.Request, in *ReqGetById) (out *ResGetById, err error) {
	command := user.GetById().WithId(in.Id)
	err = command.Execute(ctx)

	out.mapResponse(command)

	return
}

type ReqGetById struct {
	Id int64 `path:"id"`
}

type ResGetById struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
	Email  string `json:"email"`
}

func (r *ResGetById) mapResponse(command usecase.GetById) *ResGetById {
	return &ResGetById{
		Id:     command.Id(),
		Name:   command.Name(),
		Family: command.Family(),
		Email:  command.Email(),
	}
}
