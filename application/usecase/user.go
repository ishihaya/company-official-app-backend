package usecase

import (
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/repository"
	"golang.org/x/xerrors"
)

type UserUsecase interface {
	Get(authID string) (*entity.User, error)
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
	user, err := u.userRepository.GetByAuthID(authID)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return user, nil
}
