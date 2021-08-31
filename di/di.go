package di

import (
	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/interface/controller"
	"github.com/ishihaya/company-official-app-backend/interface/middleware"
	"github.com/ishihaya/company-official-app-backend/interface/operator"
	"github.com/ishihaya/company-official-app-backend/interface/repository"
	"github.com/ishihaya/company-official-app-backend/pkg/authgo"
	"github.com/ishihaya/company-official-app-backend/pkg/db"
)

func InitUser() controller.UserController {
	conn := db.New()
	repositoryUser := repository.NewUser(conn)
	operatorAppID := operator.NewAppIDOperator()
	usecaseUser := usecase.NewUser(repositoryUser, operatorAppID)
	controllerUser := controller.NewUserController(usecaseUser)
	return controllerUser
}

func InitAuth() middleware.AuthMiddleware {
	client := authgo.New()
	operatorAuth := operator.NewAuth(client)
	usecaseAuth := usecase.NewAuth(operatorAuth)
	middlewareAuth := middleware.NewAuthMiddleware(usecaseAuth)
	return middlewareAuth
}
