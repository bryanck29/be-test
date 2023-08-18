package request

type PostLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type PostRefreshLogin struct {
	Token string `json:"token" validate:"required"`
}
