package usecase

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseGetTradeConfig struct {
	TradeConfigRepo mysql.TradeConfigRepositoryInterface
}

type UsecaseGetTradeConfigInterface interface {
	GetTradeConfig(in usecasedto.InputTradeConfigDto) (out *usecasedto.OutputTradeConfigDto, err error)
}

func NewGetTradeConfigUsecase(tradeConfigRepo mysql.TradeConfigRepositoryInterface) UsecaseGetTradeConfigInterface {
	return &usecaseGetTradeConfig{
		TradeConfigRepo: tradeConfigRepo,
	}
}

func (u *usecaseGetTradeConfig) GetTradeConfig(in usecasedto.InputTradeConfigDto) (out *usecasedto.OutputTradeConfigDto, err error) {
	response, err := u.TradeConfigRepo.Get(repositorydto.InputTradeConfigDto{
		User:            in.User,
		Modality:        in.Modality,
		Strategy:        in.Strategy,
		StrategyVariant: in.StrategyVariant,
		Parity:          in.Parity,
		Exchange:        in.Exchange,
	})
	if err != nil {
		log.Panicln("GTC 00: ", err)
	}

	out = &usecasedto.OutputTradeConfigDto{}

	out.TradeConfig = response.TradeConfig
	out.User = response.User
	out.Modality = response.Modality
	out.Strategy = response.Strategy
	out.StrategyEnabled = response.StrategyEnabled
	out.StrategyVariant = response.StrategyVariant
	out.Parity = response.Parity
	out.Exchange = response.Exchange
	out.OperationQuantity = response.OperationQuantity
	out.OperationAmount = response.OperationAmount
	out.Enabled = response.Enabled
	out.DefaultProfitPercentage = response.DefaultProfitPercentage
	out.WalletValueLimit = response.WalletValueLimit
	out.ModalityName = response.ModalityName
	out.StrategyName = response.StrategyName
	out.StrategyVariantName = response.StrategyVariantName
	out.StrategyVariantEnabled = response.StrategyVariantEnabled
	out.ParitySymbol = response.ParitySymbol
	out.ExchangeName = response.ExchangeName

	return
}
