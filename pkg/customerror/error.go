package customerror

import "golang.org/x/xerrors"

var (
	ErrInternalServerError = xerrors.New("internal_server_error")
	ErrUserNotFound        = xerrors.New("user_not_found")
	ErrGetAuthID           = xerrors.New("err_get_auth_id")
)
