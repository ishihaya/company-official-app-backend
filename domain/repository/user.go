package repository

import "github.com/ishihaya/company-official-app-backend/domain/entity"

type UserRepository interface {
	GetByAuthID(authID string) (*entity.User, error)
}
