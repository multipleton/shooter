package users

import (
	"github.com/multipleton/shooter/server/fail"
)

type UsersRepository struct {
	users []*UsersEntity
}

func (ur *UsersRepository) FindById(id int) (*UsersEntity, error) {
	_, result, err := ur.findByIdWithIndex(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (ur *UsersRepository) Create(user *UsersEntity) (*UsersEntity, error) {
	id := 1
	if len(ur.users) > 0 {
		id = ur.users[len(ur.users)-1].Id + 1
	}
	user.Id = id
	// TODO: validate fields
	ur.users = append(ur.users, user)
	entity, err := ur.FindById(id)
	if err != nil {
		return nil, err
	}
	return NewUsersEntityFromUsersEntity(*entity), nil
}

func (ur *UsersRepository) Update(id int, user *UsersEntity) (*UsersEntity, error) {
	index, _, err := ur.findByIdWithIndex(id)
	if err != nil {
		return nil, err
	}
	// TODO: validate fields
	ur.users[index] = user
	entity, err := ur.FindById(id)
	if err != nil {
		return nil, err
	}
	return NewUsersEntityFromUsersEntity(*entity), nil
}

func (ur *UsersRepository) Delete(id int) error {
	index, _, err := ur.findByIdWithIndex(id)
	if err != nil {
		return err
	}
	ur.users = append(ur.users[:index], ur.users[index+1:]...)
	return nil
}

func (ur *UsersRepository) findByIdWithIndex(id int) (int, *UsersEntity, error) {
	for i, entry := range ur.users {
		if id == ur.users[i].Id {
			return i, entry, nil
		}
	}
	return -1, nil, fail.EntityNotFound(id)
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{
		users: []*UsersEntity{},
	}
}
