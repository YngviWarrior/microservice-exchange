package usecases

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecases/usecasedto"
)

type UseCasesInterface interface {
	ListExchanges() ([]*usecasedto.OutputListExchangeDto, error)
	CreateTradeConfig(*usecasedto.InputTradeConfigDto) (bool, error)
	ListTradeConfig() ([]*usecasedto.OutputTradeConfigDto, error)
}

type usecases struct {
	ExchangeRepo    mysql.ExchangeRepositoryInterface
	TradeConfigRepo mysql.TradeConfigRepositoryInterface
}

func InitUseCases(
	exchangeRepo mysql.ExchangeRepositoryInterface,
	tradeConfigRepo mysql.TradeConfigRepositoryInterface,
) UseCasesInterface {
	return &usecases{
		ExchangeRepo:    exchangeRepo,
		TradeConfigRepo: tradeConfigRepo,
	}
}
