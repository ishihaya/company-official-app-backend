package di

import (
	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/infrastructure/auth"
	"github.com/ishihaya/company-official-app-backend/infrastructure/db"
	"github.com/ishihaya/company-official-app-backend/interface/controller"
	"github.com/ishihaya/company-official-app-backend/interface/middleware"
	"github.com/ishihaya/company-official-app-backend/interface/operator"
	"github.com/ishihaya/company-official-app-backend/interface/repository"
)

func InitUser() controller.User {
	conn := db.New()
	repositoryUser := repository.NewUser(conn)
	operatorAppID := operator.NewAppID()
	usecaseUser := usecase.NewUser(repositoryUser, operatorAppID)
	controllerUser := controller.NewUser(usecaseUser)
	return controllerUser
}

func InitAuth() middleware.Auth {
	client := auth.New()
	operatorAuth := operator.NewAuth(client)
	usecaseAuth := usecase.NewAuth(operatorAuth)
	middlewareAuth := middleware.NewAuth(usecaseAuth)
	return middlewareAuth
}
