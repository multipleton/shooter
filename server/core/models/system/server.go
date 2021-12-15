package system

import "github.com/multipleton/shooter/server/http/servers"

type Server struct {
	Id    int
	Title string
	Users []*User
}

func NewServerFromServersEntity(entity servers.ServersEntity) *Server {
	var users []*User
	for _, entry := range entity.Users {
		users = append(users, NewUserFromUsersEntity(*entry))
	}
	return &Server{
		Id:    entity.Id,
		Title: entity.Title,
		Users: users,
	}
}
