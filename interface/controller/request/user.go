package request

import "time"

type UserGet struct {
	AuthID string `json:"-"`
}

type UserCreate struct {
	AuthID      string    `json:"-"`
	CurrentTime time.Time `json:"-"`
	Nickname    string    `json:"nickname" binding:"required,max=20"`
}
