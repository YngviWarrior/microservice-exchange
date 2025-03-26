package usecase

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseUpdateTradeConfig struct {
	TradeConfigRepo mysql.TradeConfigRepositoryInterface
}

type UseCaseUpdateTradeConfigInterface interface {
	UpdateTradeConfig(*usecasedto.InputUpdateTradeConfigDto) (*usecasedto.OutputUpdateTradeConfigDto, error)
}

func NewUpdateTradeConfigUseCase(tradeConfigRepo mysql.TradeConfigRepositoryInterface) UseCaseUpdateTradeConfigInterface {
	return &usecaseUpdateTradeConfig{
		TradeConfigRepo: tradeConfigRepo,
	}
}

func (u *usecaseUpdateTradeConfig) UpdateTradeConfig(in *usecasedto.InputUpdateTradeConfigDto) (out *usecasedto.OutputUpdateTradeConfigDto, err error) {
	if !u.TradeConfigRepo.Update(&repositorydto.InputTradeConfigDto{
		TradeConfig:             in.TradeConfig,
		User:                    in.User,
		Parity:                  in.Parity,
		Exchange:                in.Exchange,
		Strategy:                in.Strategy,
		StrategyVariant:         in.StrategyVariant,
		OperationQuantity:       in.OperationQuantity,
		OperationAmount:         in.OperationAmount,
		DefaultProfitPercentage: in.DefaultProfitPercentage,
		WalletValueLimit:        in.WalletValueLimit,
		Enabled:                 in.Enabled,
	}) {
		log.Println("UUTC 00: ", err)
		return
	}

	return
}
