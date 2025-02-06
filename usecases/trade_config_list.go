package usecases

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/usecases/usecasedto"
)

func (u *usecases) ListTradeConfig() (out []*usecasedto.OutputTradeConfigDto, err error) {
	response, err := u.TradeConfigRepo.List()
	if err != nil {
		log.Panicln("LTC 00: ", err)
	}

	o := &usecasedto.OutputTradeConfigDto{}
	for _, v := range response {
		o.TradeConfig = v.TradeConfig
		o.Modality = v.Modality
		o.Strategy = v.Strategy
		o.StrategyVariant = v.StrategyVariant
		o.Parity = v.Parity
		o.Exchange = v.Exchange
		o.OperationQuantity = v.OperationQuantity
		o.OperationAmount = v.OperationAmount
		o.DefaultProfitPercentage = v.DefaultProfitPercentage
		o.WalletValueLimit = v.WalletValueLimit
		o.Enabled = v.Enabled

		out = append(out, o)
	}

	return
}
