package auth

type LoginAuthDto struct {
	Username string `json:"username"`
}

type LogoutAuthDto struct {
	Id int `json:"id"`
}
