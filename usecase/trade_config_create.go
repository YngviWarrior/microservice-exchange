package usecase

import (
	"errors"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/infra/external"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseCreateTradeConfig struct {
	TradeConfigRepo mysql.TradeConfigRepositoryInterface
	ExchangeRepo    mysql.ExchangeRepositoryInterface
	ModalityRepo    mysql.ModalityRepositoryInterface
	StrategyRepo    mysql.StrategyRepositoryInterface
}

type UsecaseCreateTradeConfigInterface interface {
	CreateTradeConfig(in *usecasedto.InputTradeConfigDto) (response bool, err error)
}

func NewCreateTradeConfigUsecase(tradeConfigRepo mysql.TradeConfigRepositoryInterface, exchangeRepo mysql.ExchangeRepositoryInterface, modalityRepo mysql.ModalityRepositoryInterface, strategyRepo mysql.StrategyRepositoryInterface) UsecaseCreateTradeConfigInterface {
	return &usecaseCreateTradeConfig{
		TradeConfigRepo: tradeConfigRepo,
		ExchangeRepo:    exchangeRepo,
		ModalityRepo:    modalityRepo,
		StrategyRepo:    strategyRepo,
	}
}

func (u *usecaseCreateTradeConfig) CreateTradeConfig(in *usecasedto.InputTradeConfigDto) (response bool, err error) {
	modality := u.ModalityRepo.List(repositorydto.InputModalityDto{})

	var exist bool
	for _, v := range modality {
		if v.Modality == in.Modality {
			exist = true
		}
	}

	if !exist {
		return false, errors.New("invalid modality")
	}

	exchange := u.ExchangeRepo.List(repositorydto.InputExchangeDto{})

	for _, v := range exchange {
		if v.Exchange == in.Exchange {
			exist = true
		}
	}

	if !exist {
		return false, errors.New("invalid exchange")
	}

	strategy := u.StrategyRepo.List(repositorydto.InputStrategyDto{})

	for _, v := range strategy {
		if v.Strategy == in.Strategy {
			exist = true
		}
	}

	if !exist {
		return false, errors.New("invalid strategy")
	}

	// valid strategy_variant
	// valid parity
	externalRequest := external.NewUserExternal()
	user := externalRequest.GetUserById(uint64(in.User))

	if len(user.User) == 0 {
		return false, errors.New("invalid user")
	}

	response = u.TradeConfigRepo.Create(&repositorydto.InputTradeConfigDto{
		User:                    in.User,
		Modality:                in.Modality,
		Strategy:                in.Strategy,
		StrategyVariant:         in.StrategyVariant,
		Parity:                  in.Parity,
		Exchange:                in.Exchange,
		OperationQuantity:       in.OperationQuantity,
		OperationAmount:         in.OperationAmount,
		DefaultProfitPercentage: in.DefaultProfitPercentage,
		WalletValueLimit:        in.WalletValueLimit,
		Enabled:                 in.Enabled,
	})

	if !response {
		log.Panicln("CTC 00: ", err)
	}

	return
}
