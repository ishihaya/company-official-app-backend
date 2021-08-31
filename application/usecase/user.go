package usecase

import (
	"time"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
	"github.com/ishihaya/company-official-app-backend/domain/repository"
	"golang.org/x/xerrors"
)

type User interface {
	Get(authID string) (*entity.User, error)
	Create(authID, nickname string, currentTime time.Time) error
}

type user struct {
	userRepository repository.User
	appIDOperator  operator.AppIDOperator
}

func NewUser(
	userRepository repository.User,
	appIDOperator operator.AppIDOperator,
) User {
	return &user{
		userRepository: userRepository,
		appIDOperator:  appIDOperator,
	}
}

func (u *user) Get(authID string) (*entity.User, error) {
	user, err := u.userRepository.FindByAuthID(authID)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return user, nil
}

func (u *user) Create(authID, nickname string, currentTime time.Time) error {
	id, err := u.appIDOperator.Generate(currentTime)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	user := &entity.User{
		ID:        id,
		AuthID:    authID,
		Nickname:  nickname,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	if err := u.userRepository.Store(user); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil
}
