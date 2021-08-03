package request

type AuthAPI struct {
	IDToken string `json:"-" binding:"required"`
}
