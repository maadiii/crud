package user

import (
	"context"

	"saba/internal/application/user"

	"github.com/maadiii/hertz/server"
)

func init() {
	server.Register(Edit)
}

// [PUT] /api/v1/users/:id 204 json
func Edit(ctx context.Context, req *server.Request, in *ReqUserEdit) (out any, err error) {
	err = user.Edit().
		WithId(in.Id).
		WithName(in.Name).
		WithFamily(in.Family).
		Execute(ctx)

	return
}

type ReqUserEdit struct {
	Id     int64  `path:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
}
