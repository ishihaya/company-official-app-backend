package handler

import "golang.org/x/xerrors"

func Health() error {
	return xerrors.New("health check error!!!")
}