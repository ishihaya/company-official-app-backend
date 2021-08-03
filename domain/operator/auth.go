package operator

import "github.com/ishihaya/company-official-app-backend/domain/entity"

type AuthOperator interface {
	FindByToken(token string) (*entity.Auth, error)
}
