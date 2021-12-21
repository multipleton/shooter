package users

import "github.com/multipleton/shooter/server/core/models/system"

type UsersEntity struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (ue *UsersEntity) ToUser() *system.User {
	return &system.User{
		Id:       ue.Id,
		Username: ue.Username,
	}
}

func NewUsersEntityFromAddUserDto(dto AddUserDto) *UsersEntity {
	return &UsersEntity{
		Id:       -1,
		Username: dto.Username,
	}
}

func NewUsersEntityFromUsersEntity(entity UsersEntity) *UsersEntity {
	return &UsersEntity{
		Id:       entity.Id,
		Username: entity.Username,
	}
}
