package di

import (
	"github.com/ishihaya/company-official-app-backend/application/usecase"
	"github.com/ishihaya/company-official-app-backend/infra/db"
	"github.com/ishihaya/company-official-app-backend/interface/dataaccess/repository"
	"github.com/ishihaya/company-official-app-backend/interface/operator"
	"github.com/ishihaya/company-official-app-backend/interface/presentation/handler"
)

func InitUser() handler.UserHandler {
	conn := db.New()
	repositoryUser := repository.NewUserRepository(conn)
	appIDOperator := operator.NewAppIDOperator()
	usecaseUser := usecase.NewUserUsecase(repositoryUser, appIDOperator)
	handlerUser := handler.NewUserHandler(usecaseUser)
	return handlerUser
}
