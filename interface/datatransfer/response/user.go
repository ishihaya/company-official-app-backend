package response

import "github.com/ishihaya/company-official-app-backend/domain/entity"

type User struct {
	NickName string `json:"nickName"`
}

func NewUser(ent *entity.User) *User {
	return &User{
		NickName: ent.NickName,
	}
}

type UserGet struct {
	*User
}
