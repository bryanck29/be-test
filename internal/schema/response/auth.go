package response

type PostLogin struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh"`
}

type PostRefreshLogin struct {
	PostLogin
}
