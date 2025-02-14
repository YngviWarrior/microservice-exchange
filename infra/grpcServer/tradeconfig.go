package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) ListTradeConfig(ctx context.Context, in *pb.ListTradeConfigRequest) (out *pb.TradeConfigResponse, err error) {
	tradeConfigRepo := mysql.NewTradeConfigRepository(g.Db)
	useCases := usecase.NewUsecase(nil, tradeConfigRepo, nil, nil)

	response, err := useCases.ListTradeConfig()
	if err != nil {
		return
	}

	out = &pb.TradeConfigResponse{
		TradeConfig: []*pb.TradeConfig{},
	}

	for _, v := range response {
		var o pb.TradeConfig
		o.TradeConfig = v.TradeConfig
		o.User = v.User
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

		out.TradeConfig = append(out.TradeConfig, &o)
	}

	return
}

func (g *grpcServer) CreateTradeConfig(ctx context.Context, in *pb.CreateTradeConfigRequest) (out *pb.TradeConfigResponse, err error) {
	tradeConfigRepo := mysql.NewTradeConfigRepository(g.Db)
	modalityRepo := mysql.NewModalityRepository(g.Db)
	exchangeRepo := mysql.NewExchangeRepository(g.Db)
	strategyRepo := mysql.NewStrategyRepository(g.Db)
	useCases := usecase.NewUsecase(exchangeRepo, tradeConfigRepo, modalityRepo, strategyRepo)

	_, err = useCases.CreateTradeConfig(&usecasedto.InputTradeConfigDto{
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
	if err != nil {
		return
	}

	list, err := useCases.ListTradeConfig()
	if err != nil {
		return
	}

	out = &pb.TradeConfigResponse{
		TradeConfig: []*pb.TradeConfig{},
	}

	for _, v := range list {
		var o pb.TradeConfig
		o.TradeConfig = v.TradeConfig
		o.User = v.User
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

		out.TradeConfig = append(out.TradeConfig, &o)
	}

	return
}
