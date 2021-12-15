package auth

import "github.com/multipleton/shooter/server/http/users"

type AuthService struct {
	usersService *users.UsersService
}

func (as *AuthService) Login(dto LoginAuthDto) (*users.UsersEntity, error) {
	return as.usersService.AddUser(users.AddUserDto{
		Username: dto.Username,
	})
}

func (as *AuthService) Logout(dto LogoutAuthDto) error {
	return as.usersService.DeleteUser(dto.Id)
}

func NewAuthService(usersService *users.UsersService) *AuthService {
	return &AuthService{
		usersService: usersService,
	}
}
