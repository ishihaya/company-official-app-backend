package di

import (
	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/infra/db"
	"github.com/ishihaya/company-official-app-backend/interface/dataaccess/repository"
	"github.com/ishihaya/company-official-app-backend/interface/presentation/handler"
)

func InitUser() handler.UserHandler {
	conn := db.New()
	repositoryUser := repository.NewUserRepository(conn)
	usecaseUser := usecase.NewUserUsecase(repositoryUser)
	handlerUser := handler.NewUserHandler(usecaseUser)
	return handlerUser
}
