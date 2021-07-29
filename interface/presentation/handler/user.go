package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishihaya/company-official-app-backend/application/customerror"
	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/infra/logger"
	"github.com/ishihaya/company-official-app-backend/interface/datatransfer/response"
	"github.com/ishihaya/company-official-app-backend/pkg/context"
)

type UserHandler interface {
	Get(c *gin.Context)
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
// @Summary ユーザー情報を取得する
// @Accept  json
// @Produce  json
// @Success 200 {object} response.UserGet
// @Failure 500 {object} string "Something wrong"
// @Router /user [get]
func (u *userHandler) Get(c *gin.Context) {
	// request
	authID, err := context.GetAuthID(c)
	if err != nil {
		logger.Logging.Warnf(": %+v", err)
		c.JSON(http.StatusBadRequest, customerror.ErrGetAuthID.Error())
		return
	}

	// usecase
	user, err := u.userUsecase.Get(authID)
	if err != nil {
		logger.Logging.Errorf("failed to get user: %+v", err)
		c.JSON(http.StatusInternalServerError, customerror.ErrInternalServerError.Error())
		return
	}

	// response
	res := &response.UserGet{
		User: response.NewUserResponse(user),
	}
	c.JSON(http.StatusOK, res)
}
