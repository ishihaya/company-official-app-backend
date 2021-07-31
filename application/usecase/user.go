package usecase

import (
	"time"

	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/repository"
	"golang.org/x/xerrors"
)

type UserUsecase interface {
	Get(authID string) (*entity.User, error)
	Create(id, authID, nickName string, now time.Time) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) Get(authID string) (*entity.User, error) {
	user, err := u.userRepository.FindByAuthID(authID)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return user, nil
}

func (u *userUsecase) Create(id, authID, nickName string, now time.Time) error {
	user := &entity.User{
		ID:        id,
		AuthID:    authID,
		NickName:  nickName,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := u.userRepository.Store(user); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil
}
