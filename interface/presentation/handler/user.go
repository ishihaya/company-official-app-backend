package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/domain/service/apperror"
	"github.com/ishihaya/company-official-app-backend/interface/datatransfer/request"
	"github.com/ishihaya/company-official-app-backend/interface/datatransfer/response"
	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
	"github.com/ishihaya/company-official-app-backend/pkg/factory"
	"github.com/ishihaya/company-official-app-backend/pkg/logging"
	"golang.org/x/xerrors"
)

type UserHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
	log         logging.Log
}

func NewUserHandler(
	userUsecase usecase.UserUsecase,
) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
		log:         logging.GetInstance(),
	}
}

// Get
// @Summary 認証情報から自分のユーザー情報を取得する
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentiation header"
// @Success 200 {object} response.UserGet
// @Failure 400 {object} string "Something wrong"
// @Failure 404 {object} string "Something wrong"
// @Failure 500 {object} string "Something wrong"
// @Router /user [get]
func (u *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(request.UserGet)
	var err error
	req.AuthID, err = contextgo.AuthID(ctx)
	if err != nil {
		u.log.Warnf(": %+v", err)
		factory.JSON(w, http.StatusBadRequest, apperror.ErrGetAuthID.Error())
		return
	}

	user, err := u.userUsecase.Get(req.AuthID)
	if err != nil {
		u.log.Warnf("failed to get user: %+v", err)
		if xerrors.Is(err, apperror.ErrUserNotFound) {
			factory.JSON(w, http.StatusNotFound, apperror.ErrUserNotFound.Error())
			return
		}
		factory.JSON(w, http.StatusInternalServerError, apperror.ErrInternalServerError.Error())
		return
	}
	res := &response.UserGet{
		User: response.NewUser(user),
	}
	factory.JSON(w, http.StatusOK, res)
}

// Create
// @Summary 認証情報とリクエスト情報からユーザーを作成する
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentiation header"
// @Param data body request.UserCreate true "request body"
// @Success 204
// @Failure 500 {object} string "Something wrong"
// @Router /user [post]
func (u *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(request.UserCreate)
	var err error
	// TODO validate
	err = json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}
	req.CurrentTime, err = contextgo.CurrentTime(ctx)
	if err != nil {
		u.log.Warnf("failed to get current time: %+v", err)
		factory.JSON(w, http.StatusBadRequest, apperror.ErrGetTime.Error())
		return
	}
	req.AuthID, err = contextgo.AuthID(ctx)
	if err != nil {
		u.log.Warnf("failed to get auth id: %+v", err)
		factory.JSON(w, http.StatusBadRequest, apperror.ErrGetAuthID.Error())
		return
	}

	if err = u.userUsecase.Create(req.AuthID, req.Nickname, req.CurrentTime); err != nil {
		u.log.Errorf("failed to get user: %+v", err)
		factory.JSON(w, http.StatusInternalServerError, apperror.ErrInternalServerError.Error())
		return
	}

	factory.JSON(w, http.StatusNoContent, nil)
}
