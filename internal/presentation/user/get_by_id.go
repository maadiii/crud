package user

import (
	"context"
	"saba/internal/application/user"

	"github.com/maadiii/hertz/server"
)

func init() {
	server.Register(GetById)
}

// [GET] /api/v1/users/:id 200 json
func GetById(ctx context.Context, req *server.Request, in *ReqGetById) (out *ResGetById, err error) {
	command := user.GetById().WithId(in.Id)
	err = command.Execute(ctx)

	out.setId(in.Id).
		setName(command.Name()).
		setFamily(command.Family()).
		withEmail(command.Email())

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

func (r *ResGetById) setId(value int64) *ResGetById {
	r.Id = value

	return r
}

func (r *ResGetById) setName(value string) *ResGetById {
	r.Name = value

	return r
}

func (r *ResGetById) setFamily(value string) *ResGetById {
	r.Family = value

	return r
}

func (r *ResGetById) withEmail(value string) *ResGetById {
	r.Email = value

	return r
}
