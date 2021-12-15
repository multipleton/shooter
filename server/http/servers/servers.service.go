package servers

import "github.com/multipleton/shooter/server/http/users"

type ServersService struct {
	serversRepository *ServersRepository
	usersService      *users.UsersService
}

func (ss *ServersService) All() []*ServersEntity {
	return ss.serversRepository.ReadAll()
}

func (ss *ServersService) CreateUserServer(userId int) (*ServersEntity, error) {
	user, err := ss.usersService.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	entity := NewServersEntityFromUsersEntity(*user)
	result, err := ss.serversRepository.Create(entity)
	return result, err
}

func (ss *ServersService) JoinUserServer(serverId, userId int) error {
	user, err := ss.usersService.FindUserById(userId)
	if err != nil {
		return err
	}
	entity, err := ss.serversRepository.FindById(serverId)
	if err != nil {
		return err
	}
	entity.AddUser(*user)
	_, err = ss.serversRepository.Update(serverId, entity)
	return err
}

func (ss *ServersService) LeaveUserServer(serverId, userId int) error {
	user, err := ss.usersService.FindUserById(userId)
	if err != nil {
		return err
	}
	entity, err := ss.serversRepository.FindById(serverId)
	if err != nil {
		return err
	}
	entity.RemoveUser(*user)
	_, err = ss.serversRepository.Update(serverId, entity)
	return err
}

func NewServersService(
	serversRepository *ServersRepository,
	usersService *users.UsersService,
) *ServersService {
	return &ServersService{
		serversRepository: serversRepository,
		usersService:      usersService,
	}
}
