package usecase

import (
	"time"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/operator"
	"github.com/ishihaya/company-official-app-backend/domain/repository"
	"golang.org/x/xerrors"
)

type UserUsecase interface {
	Get(authID string) (*entity.User, error)
	Create(authID, nickname string, currentTime time.Time) error
}

type userUsecase struct {
	userRepository repository.UserRepository
	appIDOperator  operator.AppIDOperator
}

func NewUserUsecase(
	userRepository repository.UserRepository,
	appIDOperator operator.AppIDOperator,
) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		appIDOperator:  appIDOperator,
	}
}

func (u *userUsecase) Get(authID string) (*entity.User, error) {
	user, err := u.userRepository.FindByAuthID(authID)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return user, nil
}

func (u *userUsecase) Create(authID, nickname string, currentTime time.Time) error {
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
