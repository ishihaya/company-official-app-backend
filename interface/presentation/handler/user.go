package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishihaya/company-official-app-backend/application/customerror"
	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/infra/logger"
	"github.com/ishihaya/company-official-app-backend/interface/datatransfer/response"
	"github.com/ishihaya/company-official-app-backend/pkg/contextgo"
	"github.com/ishihaya/company-official-app-backend/pkg/ulidgo"
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
	authID, err := contextgo.GetAuthID(c)
	if err != nil {
		logger.Logging.Warnf(": %+v", err)
		c.JSON(http.StatusBadRequest, customerror.ErrGetAuthID.Error())
		return
	}

	// usecase
	user, err := u.userUsecase.Get(authID)
	if err != nil {
		logger.Logging.Warnf("failed to get user: %+v", err)
		if xerrors.Is(err, customerror.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, customerror.ErrUserNotFound.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, customerror.ErrInternalServerError.Error())
		return
	}

	// response
	res := &response.UserGet{
		User: response.NewUserResponse(user),
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
	now, err := contextgo.Now(c)
	if err != nil {
		// TODO
		return
	}
	ulid, err := ulidgo.Generate(now)
	if err != nil {
		// TODO
		return
	}
	authID, err := contextgo.GetAuthID(c)
	if err != nil {
		// TODO
		// 	logger.Logging.Warnf(": %+v", err)
		// 	c.JSON(http.StatusBadRequest, customerror.ErrGetAuthID.Error())
		return
	}
	// TODO
	// nick_name

	// if err := u.userUsecase.Create()
	// TODO
	// Response関数が必要かどうかRequestは使用していないのにどうなのか考える
}
