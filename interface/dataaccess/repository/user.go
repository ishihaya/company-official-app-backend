package repository

import (
	"database/sql"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/repository"
	"github.com/ishihaya/company-official-app-backend/infra/db"
	"github.com/ishihaya/company-official-app-backend/interface/dataaccess/dao"
	"golang.org/x/xerrors"
)

type userRepository struct {
	conn *db.Conn
}

func NewUserRepository(conn *db.Conn) repository.UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (u *userRepository) FindByAuthID(authID string) (*entity.User, error) {
	daoUser := new(dao.User)
	if err := u.conn.Get(daoUser, "SELECT * FROM users WHERE auth_id = ?", authID); err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			return nil, xerrors.Errorf("user not found authID= %s : %w", authID, entity.ErrUserNotFound)
		}
		return nil, xerrors.Errorf("failed to get user by authID= %s : %w", authID, err)
	}
	user := daoUser.ConvertToEntity()
	return user, nil
}

func (u *userRepository) Store(user *entity.User) error {
	// TODO
	return nil
}
