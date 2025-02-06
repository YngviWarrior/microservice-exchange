package usecases

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecases/usecasedto"
)

func (u *usecases) CreateTradeConfig(in *usecasedto.InputTradeConfigDto) (response bool, err error) {
	response = u.TradeConfigRepo.Create(&repositorydto.InputTradeConfigDto{
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
