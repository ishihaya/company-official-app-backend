package apperror

import (
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"golang.org/x/xerrors"
)

var (
	ErrInternalServerError entity.AppError = xerrors.New("internal_server_error")
	ErrUserNotFound        entity.AppError = xerrors.New("user_not_found")
	ErrValidation          entity.AppError = xerrors.New("validation_error")
	ErrGetAuthID           entity.AppError = xerrors.New("err_get_auth_id")
	ErrGetTime             entity.AppError = xerrors.New("err_get_time")
)
