package repository

import (
	"database/sql"
	"time"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/repository"
	"github.com/ishihaya/company-official-app-backend/domain/service/apperror"
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

type daoUser struct {
	ID        entity.AppID `db:"id"`
	AuthID    string       `db:"auth_id"`
	Nickname  string       `db:"nick_name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}

func (du *daoUser) convertUserDAOToEntity() *entity.User {
	return &entity.User{
		ID:        du.ID,
		AuthID:    du.AuthID,
		Nickname:  du.Nickname,
		CreatedAt: du.CreatedAt,
		UpdatedAt: du.UpdatedAt,
	}
}

func convertUserEntityToDAO(ent *entity.User) *daoUser {
	return &daoUser{
		ID:        ent.ID,
		AuthID:    ent.AuthID,
		Nickname:  ent.Nickname,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}

func (u *userRepository) FindByAuthID(authID string) (*entity.User, error) {
	du := new(daoUser)
	if err := u.conn.Get(du, "SELECT * FROM users WHERE auth_id = ?", authID); err != nil {
		if xerrors.Is(err, sql.ErrNoRows) {
			return nil, xerrors.Errorf("user not found authID= %s : %w", authID, apperror.ErrUserNotFound)
		}
		return nil, xerrors.Errorf("failed to get user by authID= %s : %w", authID, err)
	}
	return du.convertUserDAOToEntity(), nil
}

func (u *userRepository) Store(user *entity.User) error {
	du := convertUserEntityToDAO(user)
	if _, err := u.conn.NamedExec("INSERT INTO users(id, auth_id, nick_name, created_at, updated_at) VALUES(:id, :auth_id, :nick_name, :created_at, :updated_at)", du); err != nil {
		return xerrors.Errorf("failed to create user : %w", err)
	}
	return nil
}
