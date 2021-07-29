package customerror

import "golang.org/x/xerrors"

var (
	ErrGetAuthID = xerrors.New("failed to get authID")
)
