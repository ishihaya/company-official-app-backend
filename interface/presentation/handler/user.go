package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/domain/entity"
	"github.com/ishihaya/company-official-app-backend/interface/datatransfer/request"
	"github.com/ishihaya/company-official-app-backend/interface/datatransfer/response"
	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
	"github.com/ishihaya/company-official-app-backend/pkg/logger"
	"golang.org/x/xerrors"
)

type UserHandler interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

// Get
// @Summary 認証情報から自分のユーザー情報を取得する
// @Accept  json
// @Produce  json
// @Success 200 {object} response.UserGet
// @Failure 400 {object} string "Something wrong"
// @Failure 404 {object} string "Something wrong"
// @Failure 500 {object} string "Something wrong"
// @Router /user [get]
func (u *userHandler) Get(c *gin.Context) {
	// request
	req := new(request.UserGet)
	var err error
	req.AuthID, err = contextgo.GetAuthID(c)
	if err != nil {
		logger.Logging.Warnf(": %+v", err)
		c.JSON(http.StatusBadRequest, entity.ErrGetAuthID.Error())
		return
	}

	// usecase
	user, err := u.userUsecase.Get(req.AuthID)
	if err != nil {
		logger.Logging.Warnf("failed to get user: %+v", err)
		if xerrors.Is(err, entity.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, entity.ErrUserNotFound.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, entity.ErrInternalServerError.Error())
		return
	}

	// response
	res := &response.UserGet{
		User: response.NewUser(user),
	}
	c.JSON(http.StatusOK, res)
}

// Create
// @Summary 認証情報とリクエスト情報からユーザーを作成する
// @Accept  json
// @Produce  json
// @Success 204
// @Failure 500 {object} string "Something wrong"
// @Router /user [post]
func (u *userHandler) Create(c *gin.Context) {
	req := new(request.UserCreate)
	var err error
	if err = c.ShouldBindJSON(req); err != nil {
		// TODO
		return
	}
	req.CurrentTime, err = contextgo.Now(c)
	if err != nil {
		// TODO
		return
	}
	req.AuthID, err = contextgo.GetAuthID(c)
	if err != nil {
		// TODO
		// 	logger.Logging.Warnf(": %+v", err)
		// 	c.JSON(http.StatusBadRequest, entity.ErrGetAuthID.Error())
		return
	}

	if err := u.userUsecase.Create(req.AuthID, req.NickName, req.CurrentTime); err != nil {
		// TODO
		return
	}
	// TODO
}
