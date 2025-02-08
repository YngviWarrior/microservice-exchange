package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type UseCaseInterface interface {
	ListExchanges() ([]*usecasedto.OutputListExchangeDto, error)
	CreateTradeConfig(*usecasedto.InputTradeConfigDto) (bool, error)
	ListTradeConfig() ([]*usecasedto.OutputTradeConfigDto, error)
}

// CreateTradeConfig implements usecases.UseCasesInterface.

type usecase struct {
	ExchangeRepo    mysql.ExchangeRepositoryInterface
	TradeConfigRepo mysql.TradeConfigRepositoryInterface
}

func NewUsecase(
	exchangeRepo mysql.ExchangeRepositoryInterface,
	tradeConfigRepo mysql.TradeConfigRepositoryInterface,
) UseCaseInterface {
	return &usecase{
		ExchangeRepo:    exchangeRepo,
		TradeConfigRepo: tradeConfigRepo,
	}
}
