package repository

import (
	"database/sql"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/repository"
	"github.com/ishihaya/company-official-app-backend/domain/service/apperror"
	"github.com/ishihaya/company-official-app-backend/interface/dataaccess/dao"
	"github.com/ishihaya/company-official-app-backend/pkg/db"
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
			return nil, xerrors.Errorf("user not found authID= %s : %w", authID, apperror.ErrUserNotFound)
		}
		return nil, xerrors.Errorf("failed to get user by authID= %s : %w", authID, err)
	}
	user := daoUser.ConvertToEntity()
	return user, nil
}

func (u *userRepository) Store(user *entity.User) error {
	daoUser := dao.ConvertToDAOUser(user)
	if _, err := u.conn.NamedExec("INSERT INTO users(id, auth_id, nick_name, created_at, updated_at) VALUES(:id, :auth_id, :nick_name, :created_at, :updated_at)", daoUser); err != nil {
		return xerrors.Errorf("failed to create user : %w", err)
	}
	return nil
}
