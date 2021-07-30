package customerror

import "golang.org/x/xerrors"

var (
	ErrUserNotFound = xerrors.New("user_not_found")
)
