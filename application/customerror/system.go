package customerror

import "golang.org/x/xerrors"

var (
	ErrInternalServerError = xerrors.New("internal_server_error")
)
