package dao

import (
	"time"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
)

type User struct {
	ID        string    `db:"id"`
	AuthID    string    `db:"auth_id"`
	NickName  string    `db:"nick_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) ConvertToEntity() *entity.User {
	return &entity.User{
		ID:        u.ID,
		AuthID:    u.AuthID,
		NickName:  u.NickName,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
