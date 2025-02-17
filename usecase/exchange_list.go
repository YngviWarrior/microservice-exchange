package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListExchange struct {
	TradeConfigRepo mysql.TradeConfigRepositoryInterface
	ExchangeRepo    mysql.ExchangeRepositoryInterface
	ModalityRepo    mysql.ModalityRepositoryInterface
	StrategyRepo    mysql.StrategyRepositoryInterface
}

type UsecaseListExchangeInterface interface {
	ListExchange() ([]*usecasedto.OutputListExchangeDto, error)
}

func NewListExchangeUsecase(exchangeRepo mysql.ExchangeRepositoryInterface) UsecaseListExchangeInterface {
	return &usecaseListExchange{
		ExchangeRepo: exchangeRepo,
	}
}

func (u *usecaseListExchange) ListExchange() ([]*usecasedto.OutputListExchangeDto, error) {
	response := []*usecasedto.OutputListExchangeDto{}
	list := u.ExchangeRepo.List(repositorydto.InputExchangeDto{})

	for _, v := range list {
		x := &usecasedto.OutputListExchangeDto{}

		x.Exchange = v.Exchange
		x.Name = v.Name

		response = append(response, x)
	}

	return response, nil
}
