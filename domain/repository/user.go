package repository

import "github.com/ishihaya/company-official-app-backend/domain/entity"

type User interface {
	FindByAuthID(authID string) (*entity.User, error)
	Store(user *entity.User) error
}
