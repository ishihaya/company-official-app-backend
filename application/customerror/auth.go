package customerror

import "golang.org/x/xerrors"

var (
	ErrGetAuthID = xerrors.New("err_get_auth_id")
)
