package entity

import "golang.org/x/xerrors"

// AppError - アプリケーションで使用されるエラーのオブジェクト
type AppError error

var (
	ErrInternalServerError AppError = xerrors.New("internal_server_error")
	ErrUserNotFound        AppError = xerrors.New("user_not_found")
	ErrGetAuthID           AppError = xerrors.New("err_get_auth_id")
)
