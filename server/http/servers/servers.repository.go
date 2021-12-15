package servers

import "github.com/multipleton/shooter/server/fail"

type ServersRepository struct {
	servers []*ServersEntity
}

func (sr *ServersRepository) FindById(id int) (*ServersEntity, error) {
	_, result, err := sr.findByIdWithIndex(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sr *ServersRepository) Create(server *ServersEntity) (*ServersEntity, error) {
	id := 1
	if len(sr.servers) > 0 {
		id = sr.servers[len(sr.servers)-1].Id + 1
	}
	server.Id = id
	// TODO: validate fields
	sr.servers = append(sr.servers, server)
	entity, err := sr.FindById(id)
	if err != nil {
		return nil, err
	}
	return NewServersEntityFromServersEntity(*entity), nil
}

func (sr *ServersRepository) Update(id int, server *ServersEntity) (*ServersEntity, error) {
	index, _, err := sr.findByIdWithIndex(id)
	if err != nil {
		return nil, err
	}
	// TODO: validate fields
	sr.servers[index] = server
	entity, err := sr.FindById(id)
	if err != nil {
		return nil, err
	}
	return NewServersEntityFromServersEntity(*entity), nil
}

func (sr *ServersRepository) Delete(id int) error {
	index, _, err := sr.findByIdWithIndex(id)
	if err != nil {
		return err
	}
	sr.servers = append(sr.servers[:index], sr.servers[index+1:]...)
	return nil
}

func (sr *ServersRepository) ReadAll() []*ServersEntity {
	return sr.servers
}

func (sr *ServersRepository) findByIdWithIndex(id int) (int, *ServersEntity, error) {
	for i, entry := range sr.servers {
		if id == sr.servers[i].Id {
			return i, entry, nil
		}
	}
	return -1, nil, fail.EntityNotFound(id)
}

func NewServersRepository() *ServersRepository {
	return &ServersRepository{
		servers: []*ServersEntity{},
	}
}
