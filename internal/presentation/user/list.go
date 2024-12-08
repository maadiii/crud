package user

import (
	"context"

	"github.com/maadiii/crud/internal/application/user"

	"github.com/maadiii/hertz/server"
)

func init() {
	server.Register(List)
}

// [GET] /api/v1/users 200 json
func List(ctx context.Context, req *server.Request, in *RequestList) (out *ResponseList, err error) {
	command := user.List().
		SetPagination(in.From, in.Count).
		SetSearch(in.SearchKey, in.SearchValue).
		SetSort(in.Sort, in.Desc)
	err = command.Execute(ctx)

	out = new(ResponseList)
	for _, user := range command.Users() {
		out.Users = append(out.Users, User{
			Id:     user.Id,
			Name:   user.Name,
			Family: user.Family,
			Email:  user.Email,
		})
	}

	return
}

type RequestList struct {
	From        int    `query:"from"`
	Count       int    `query:"count"`
	Sort        string `query:"sort"`
	Desc        bool   `query:"desc"`
	SearchKey   string `query:"search-key"`
	SearchValue any    `query:"search-value"`
}

type ResponseList struct {
	Users []User `json:"users,omitempty"`
}
