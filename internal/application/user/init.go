package user

import (
	"sync"

	"github.com/maadiii/crud/pkg/usecase/persistence/user"
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
