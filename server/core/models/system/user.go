package system

import "github.com/multipleton/shooter/server/http/users"

type User struct {
	Id       int
	Username string
}

func NewUserFromUsersEntity(entity users.UsersEntity) *User {
	return &User{
		Id:       entity.Id,
		Username: entity.Username,
	}
}
