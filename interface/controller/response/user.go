package response

import "github.com/ishihaya/company-official-app-backend/domain/entity"

type User struct {
	Nickname string `json:"nickname"`
}

func NewUser(ent *entity.User) *User {
	return &User{
		Nickname: ent.Nickname,
	}
}

type UserGet struct {
	*User
}
