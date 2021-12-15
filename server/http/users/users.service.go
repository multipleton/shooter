package users

type UsersService struct {
	usersRepository *UsersRepository
}

func (us *UsersService) FindUserById(id int) (*UsersEntity, error) {
	return us.usersRepository.FindById(id)
}

func (us *UsersService) AddUser(dto AddUserDto) (*UsersEntity, error) {
	entity := NewUsersEntityFromAddUserDto(dto)
	result, err := us.usersRepository.Create(entity)
	return result, err
}

func (us *UsersService) DeleteUser(id int) error {
	return us.usersRepository.Delete(id)
}

func NewUsersService(usersRepository *UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}
