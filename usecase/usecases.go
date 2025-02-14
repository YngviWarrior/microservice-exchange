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
	ModalityRepo    mysql.ModalityRepositoryInterface
	StrategyRepo    mysql.StrategyRepositoryInterface
}

func NewUsecase(
	exchangeRepo mysql.ExchangeRepositoryInterface,
	tradeConfigRepo mysql.TradeConfigRepositoryInterface,
	modalityRepo mysql.ModalityRepositoryInterface,
	strategyRepo mysql.StrategyRepositoryInterface,
) UseCaseInterface {
	return &usecase{
		ExchangeRepo:    exchangeRepo,
		TradeConfigRepo: tradeConfigRepo,
		ModalityRepo:    modalityRepo,
		StrategyRepo:    strategyRepo,
	}
}
