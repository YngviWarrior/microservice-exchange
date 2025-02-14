package usecase

import (
	"errors"
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/infra/external"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (u *usecase) CreateTradeConfig(in *usecasedto.InputTradeConfigDto) (response bool, err error) {
	modalityList := u.ModalityRepo.List(repositorydto.InputModalityDto{})

	var exist bool
	for _, v := range modalityList {
		if v.Modality == in.Modality {
			exist = true
		}
	}

	if !exist {
		return false, errors.New("invalid modality")
	}

	exchangeList := u.ExchangeRepo.List(repositorydto.InputExchangeDto{})

	for _, v := range exchangeList {
		if v.Exchange == in.Exchange {
			exist = true
		}
	}

	if !exist {
		return false, errors.New("invalid exchange")
	}

	repositoryList := u.StrategyRepo.List(repositorydto.InputStrategyDto{})

	for _, v := range repositoryList {
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
