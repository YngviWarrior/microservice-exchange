package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/infra/internal"
	"github.com/YngviWarrior/microservice-exchange/infra/internal/internaldto"
	"github.com/YngviWarrior/microservice-exchange/utils"
)

func (g *grpcServer) Strategy(ctx context.Context, in *pb.StrategyRequest) (out *pb.StrategyResponse, err error) {
	parityRepo := mysql.NewParityRepository(g.Db)
	operationRepo := mysql.NewOperationRepository(g.Db)
	operationHistoryRepo := mysql.NewOperationHistoryRepository(g.Db)
	walletRepo := mysql.NewWalletRepository(g.Db)
	averageRepo := mysql.NewAveragePriceRepository(g.Db)

	switch in.GetStrategy() {
	case "Day":
		usecase := internal.NewUsecaseStrategy(operationRepo, operationHistoryRepo, walletRepo, parityRepo, averageRepo)

		_, err = usecase.AveragePrice(&internaldto.InputStrategyAveragePriceDto{
			ClosePrice: utils.ParseFloat(in.GetKline().GetClose()),
			TradeConfig: &internaldto.TradeConfig{
				TradeConfig:             in.GetTradeConfig().GetTradeConfig(),
				User:                    in.GetTradeConfig().GetUser(),
				Modality:                in.GetTradeConfig().GetModality(),
				Strategy:                in.GetTradeConfig().GetStrategy(),
				StrategyEnabled:         in.GetTradeConfig().GetStrategyEnabled(),
				StrategyVariant:         in.GetTradeConfig().GetStrategyVariant(),
				Parity:                  in.GetTradeConfig().GetParity(),
				Exchange:                in.GetTradeConfig().GetExchange(),
				OperationQuantity:       in.GetTradeConfig().GetOperationQuantity(),
				OperationAmount:         in.GetTradeConfig().GetOperationAmount(),
				Enabled:                 in.GetTradeConfig().GetEnabled(),
				DefaultProfitPercentage: in.GetTradeConfig().GetDefaultProfitPercentage(),
				WalletValueLimit:        in.GetTradeConfig().GetWalletValueLimit(),
				UserName:                in.GetTradeConfig().GetUserName(),
				ModalityName:            in.GetTradeConfig().GetModalityName(),
				StrategyName:            in.GetTradeConfig().GetStrategyName(),
				StrategyVariantName:     in.GetTradeConfig().GetStrategyVariantName(),
				StrategyVariantEnabled:  in.GetTradeConfig().GetStrategyVariantEnabled(),
				SymbolName:              in.GetTradeConfig().GetSymbolName(),
				ExchangeName:            in.GetTradeConfig().GetExchangeName(),
				ParitySymbol:            in.GetTradeConfig().GetParitySymbol(),
			},
		})
		if err != nil {
			return
		}

	case "OpenClose":
		usecase := internal.NewUsecaseStrategy(operationRepo, operationHistoryRepo, walletRepo, parityRepo, averageRepo)

		_, err = usecase.OpenClose(&internaldto.InputStrategyOpenCloseDto{
			ClosePrice: utils.ParseFloat(in.GetKline().GetClose()),
			TradeConfig: &internaldto.TradeConfig{
				TradeConfig:             in.GetTradeConfig().GetTradeConfig(),
				User:                    in.GetTradeConfig().GetUser(),
				Modality:                in.GetTradeConfig().GetModality(),
				Strategy:                in.GetTradeConfig().GetStrategy(),
				StrategyEnabled:         in.GetTradeConfig().GetStrategyEnabled(),
				StrategyVariant:         in.GetTradeConfig().GetStrategyVariant(),
				Parity:                  in.GetTradeConfig().GetParity(),
				Exchange:                in.GetTradeConfig().GetExchange(),
				OperationQuantity:       in.GetTradeConfig().GetOperationQuantity(),
				OperationAmount:         in.GetTradeConfig().GetOperationAmount(),
				Enabled:                 in.GetTradeConfig().GetEnabled(),
				DefaultProfitPercentage: in.GetTradeConfig().GetDefaultProfitPercentage(),
				WalletValueLimit:        in.GetTradeConfig().GetWalletValueLimit(),
				UserName:                in.GetTradeConfig().GetUserName(),
				ModalityName:            in.GetTradeConfig().GetModalityName(),
				StrategyName:            in.GetTradeConfig().GetStrategyName(),
				StrategyVariantName:     in.GetTradeConfig().GetStrategyVariantName(),
				StrategyVariantEnabled:  in.GetTradeConfig().GetStrategyVariantEnabled(),
				SymbolName:              in.GetTradeConfig().GetSymbolName(),
				ExchangeName:            in.GetTradeConfig().GetExchangeName(),
				ParitySymbol:            in.GetTradeConfig().GetParitySymbol(),
			},
		})
		if err != nil {
			return
		}

	}

	return
}
