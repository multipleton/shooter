package users

type UsersEntity struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
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
