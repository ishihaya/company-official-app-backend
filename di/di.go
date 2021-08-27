package di

import (
	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/interface/dataaccess/repository"
	"github.com/ishihaya/company-official-app-backend/interface/operator"
	"github.com/ishihaya/company-official-app-backend/interface/presentation/handler"
	"github.com/ishihaya/company-official-app-backend/interface/presentation/middleware"
	"github.com/ishihaya/company-official-app-backend/pkg/authgo"
	"github.com/ishihaya/company-official-app-backend/pkg/db"
	"github.com/ishihaya/company-official-app-backend/pkg/logging"
)

func InitUser() handler.UserHandler {
	conn := db.GetInstance()
	repositoryUser := repository.NewUserRepository(conn)
	operatorAppID := operator.NewAppIDOperator()
	usecaseUser := usecase.NewUserUsecase(repositoryUser, operatorAppID)
	logging := logging.GetInstance()
	handlerUser := handler.NewUserHandler(usecaseUser, logging)
	return handlerUser
}

func InitAuth() middleware.AuthMiddleware {
	client := authgo.New()
	operatorAuth := operator.NewAuthOperator(client)
	usecaseAuth := usecase.NewAuthUsecase(operatorAuth)
	logging := logging.GetInstance()
	middlewareAuth := middleware.NewAuthMiddleware(usecaseAuth, logging)
	return middlewareAuth
}
