package user

type User struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Family string `json:"family,omitempty"`
	Email  string `json:"email,omitempty"`
}

type UserList struct {
	Users []User `json:"users,omitempty"`
}
