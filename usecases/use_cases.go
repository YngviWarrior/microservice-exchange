package usecases

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecases/usecasedto"
)

type UseCasesInterface interface {
	ListExchanges() ([]*usecasedto.ListExchangeDto, error)
}

type usecases struct {
	ExchangeRepo mysql.ExchangeRepositoryInterface
}

func InitUseCases(
	exchangeRepo mysql.ExchangeRepositoryInterface,
) UseCasesInterface {
	return &usecases{
		ExchangeRepo: exchangeRepo,
	}
}
