package entity

import "github.com/maadiii/regorm"

type User struct {
	Id       int64
	Name     string
	Family   string
	Email    string
	Password string

	*regorm.Entity[*User]
}

func (User) TableName() string {
	return "users"
}
