package usecase

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListTradeConfig struct {
	TradeConfigRepo mysql.TradeConfigRepositoryInterface
}

type UsecaseListTradeConfigInterface interface {
	ListTradeConfig() (out []*usecasedto.OutputTradeConfigDto, err error)
}

func NewListTradeConfigUsecase(tradeConfigRepo mysql.TradeConfigRepositoryInterface) UsecaseListTradeConfigInterface {
	return &usecaseListTradeConfig{
		TradeConfigRepo: tradeConfigRepo,
	}
}

func (u *usecaseListTradeConfig) ListTradeConfig() (out []*usecasedto.OutputTradeConfigDto, err error) {
	response, err := u.TradeConfigRepo.List()
	if err != nil {
		log.Panicln("LTC 00: ", err)
	}

	o := &usecasedto.OutputTradeConfigDto{}
	for _, v := range response {
		o.TradeConfig = v.TradeConfig
		o.User = v.User
		o.Modality = v.Modality
		o.Strategy = v.Strategy
		o.StrategyEnabled = v.StrategyEnabled
		o.StrategyVariant = v.StrategyVariant
		o.Parity = v.Parity
		o.Exchange = v.Exchange
		o.OperationQuantity = v.OperationQuantity
		o.OperationAmount = v.OperationAmount
		o.Enabled = v.Enabled
		o.DefaultProfitPercentage = v.DefaultProfitPercentage
		o.WalletValueLimit = v.WalletValueLimit
		o.ModalityName = v.ModalityName
		o.StrategyName = v.StrategyName
		o.StrategyVariantName = v.StrategyVariantName
		o.StrategyVariantEnabled = v.StrategyVariantEnabled
		o.ParitySymbol = v.ParitySymbol
		o.ExchangeName = v.ExchangeName

		out = append(out, o)
	}

	return
}
