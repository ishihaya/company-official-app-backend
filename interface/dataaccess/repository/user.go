package repository

import (
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/repository"
	"github.com/ishihaya/company-official-app-backend/infra/db"
)

type userRepository struct {
	conn *db.Conn
}

func NewUserRepository(conn *db.Conn) repository.UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (u *userRepository) GetByAuthID(authID string) (*entity.User, error) {
	// TODO
	return nil, nil
}
