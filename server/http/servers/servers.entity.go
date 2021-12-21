package servers

import (
	"fmt"

	"github.com/multipleton/shooter/server/core/models/system"
	"github.com/multipleton/shooter/server/http/users"
)

type ServersEntity struct {
	Id    int                  `json:"id"`
	Title string               `json:"title"`
	Users []*users.UsersEntity `json:"users"`
}

func (se *ServersEntity) AddUser(user users.UsersEntity) {
	se.Users = append(se.Users, &user)
}

func (se *ServersEntity) RemoveUser(user users.UsersEntity) {
	for i, entry := range se.Users {
		if entry.Id == user.Id {
			se.Users = append(se.Users[:i], se.Users[i+1:]...)
			return
		}
	}
}

func (se *ServersEntity) ToServer() *system.Server {
	var users []*system.User
	for _, user := range se.Users {
		users = append(users, user.ToUser())
	}
	return &system.Server{
		Id:    se.Id,
		Title: se.Title,
		Users: users,
	}
}

func NewServersEntityFromServersEntity(entity ServersEntity) *ServersEntity {
	var copiedUsers []*users.UsersEntity
	for _, user := range entity.Users {
		copiedUser := users.NewUsersEntityFromUsersEntity(*user)
		copiedUsers = append(copiedUsers, copiedUser)
	}
	return &ServersEntity{
		Id:    entity.Id,
		Title: entity.Title,
		Users: copiedUsers,
	}
}

func NewServersEntityFromUsersEntity(entity users.UsersEntity) *ServersEntity {
	return &ServersEntity{
		Id:    -1,
		Title: fmt.Sprintf("%s's server", entity.Username),
		Users: []*users.UsersEntity{&entity},
	}
}
