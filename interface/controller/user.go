package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/domain/service/apperror"
	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
	"github.com/ishihaya/company-official-app-backend/pkg/factory"
	"github.com/ishihaya/company-official-app-backend/pkg/logging"
	"golang.org/x/xerrors"
)

type User interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type user struct {
	userUsecase usecase.User
	log         logging.Log
}

func NewUser(
	userUsecase usecase.User,
) User {
	return &user{
		userUsecase: userUsecase,
		log:         logging.GetInstance(),
	}
}

// outputUser is used response
type outputUser struct {
	ID        entity.AppID `json:"id"`
	AuthID    string       `json:"authID"`
	Nickname  string       `json:"nickname"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"ureatedAt"`
}

func convertUserEntityToOutput(ent *entity.User) *outputUser {
	return &outputUser{
		ID:        ent.ID,
		AuthID:    ent.AuthID,
		Nickname:  ent.Nickname,
		CreatedAt: ent.CreatedAt,
		UpdatedAt: ent.UpdatedAt,
	}
}

type userGetRes struct {
	*outputUser
}

// Get
// @Summary 認証情報から自分のユーザー情報を取得する
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentiation header"
// @Failure 400 {object} string "Something wrong"
// @Failure 404 {object} string "Something wrong"
// @Failure 500 {object} string "Something wrong"
// @Router /user [get]
func (u *user) Get(w http.ResponseWriter, r *http.Request) {
	authID, err := contextgo.AuthID(r.Context())
	if err != nil {
		u.log.Warnf(": %+v", err)
		factory.JSON(w, http.StatusBadRequest, apperror.ErrGetAuthID.Error())
		return
	}

	user, err := u.userUsecase.Get(authID)
	if err != nil {
		u.log.Warnf("failed to get user: %+v", err)
		if xerrors.Is(err, apperror.ErrUserNotFound) {
			factory.JSON(w, http.StatusNotFound, apperror.ErrUserNotFound.Error())
			return
		}
		factory.JSON(w, http.StatusInternalServerError, apperror.ErrInternalServerError.Error())
		return
	}

	res := &userGetRes{convertUserEntityToOutput(user)}
	factory.JSON(w, http.StatusOK, res)
}

type userCreateReq struct {
	Nickname string `json:"nickname" binding:"required"`
}

// Create
// @Summary 認証情報とリクエスト情報からユーザーを作成する
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authentiation header"
// @Param data body userCreateReq true "request body"
// @Success 204
// @Failure 500 {object} string "Something wrong"
// @Router /user [post]
func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(userCreateReq)
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		panic(err)
	}
	currentTime, err := contextgo.CurrentTime(ctx)
	if err != nil {
		u.log.Warnf("failed to get current time: %+v", err)
		factory.JSON(w, http.StatusBadRequest, apperror.ErrGetTime.Error())
		return
	}
	authID, err := contextgo.AuthID(ctx)
	if err != nil {
		u.log.Warnf("failed to get auth id: %+v", err)
		factory.JSON(w, http.StatusBadRequest, apperror.ErrGetAuthID.Error())
		return
	}

	if err := u.userUsecase.Create(authID, req.Nickname, currentTime); err != nil {
		u.log.Errorf("failed to get user: %+v", err)
		factory.JSON(w, http.StatusInternalServerError, apperror.ErrInternalServerError.Error())
		return
	}

	factory.JSON(w, http.StatusNoContent, nil)
}
