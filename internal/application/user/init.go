package user

import (
	"saba/pkg/usecase/persistence/user"
	"sync"
)

func Inject(data user.Persistence) {
	once.Do(func() {
		users = data
	})
}

var (
	users user.Persistence
	once  sync.Once
)
