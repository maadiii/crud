package user

import (
	"context"
	"saba/pkg/entity"
	"saba/pkg/usecase/application/user"

	. "github.com/maadiii/regorm"
)

func List() user.List {
	return new(list)
}

type list struct {
	from, count int
	sort        string
	desc        bool
	searchKey   string
	searchValue any

	users []*entity.User
}

func (l *list) Execute(ctx context.Context) (err error) {
	stmt := SQL(new(entity.User)).Select("*")

	if len(l.searchKey) > 0 {
		stmt = stmt.Where(EQ(l.searchKey, l.searchValue))
	}

	if len(l.sort) > 0 {
		stmt = stmt.OrderBy(l.sort, l.desc)
	}

	l.users, err = stmt.Offset(l.from).Limit(l.count).Find(ctx)

	return
}

func (l *list) SetPagination(from int, count int) user.List {
	l.from, l.count = from, count

	return l
}

func (l *list) SetSort(value string, desc bool) user.List {
	l.sort = value
	l.desc = desc

	return l
}

func (l *list) SetSearch(key string, value any) user.List {
	l.searchKey = key
	l.searchValue = value

	return l
}

func (l *list) Users() []*entity.User {
	return l.users
}
