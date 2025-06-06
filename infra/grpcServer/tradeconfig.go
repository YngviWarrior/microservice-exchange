package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) GetTradeConfig(ctx context.Context, in *pb.GetTradeConfigRequest) (out *pb.GetTradeConfigResponse, err error) {
	tradeConfigRepo := mysql.NewTradeConfigRepository(g.Db)
	useCases := usecase.NewGetTradeConfigUsecase(tradeConfigRepo)

	response, err := useCases.GetTradeConfig(usecasedto.InputTradeConfigDto{
		User:            in.GetUser(),
		Modality:        in.GetModality(),
		Strategy:        in.GetStrategy(),
		StrategyVariant: in.GetStrategyVariant(),
		Parity:          in.GetParity(),
		Exchange:        in.GetExchange(),
	})
	if err != nil {
		return
	}

	out = &pb.GetTradeConfigResponse{
		TradeConfig: &pb.TradeConfig{},
	}

	var o pb.TradeConfig
	o.TradeConfig = response.TradeConfig
	o.User = response.User
	o.Modality = response.Modality
	o.Strategy = response.Strategy
	o.StrategyEnabled = response.StrategyEnabled
	o.StrategyVariant = response.StrategyVariant
	o.Parity = response.Parity
	o.Exchange = response.Exchange
	o.OperationQuantity = response.OperationQuantity
	o.OperationAmount = response.OperationAmount
	o.Enabled = response.Enabled
	o.DefaultProfitPercentage = response.DefaultProfitPercentage
	o.WalletValueLimit = response.WalletValueLimit
	o.ModalityName = response.ModalityName
	o.StrategyName = response.StrategyName
	o.StrategyVariantName = response.StrategyVariantName
	o.StrategyVariantEnabled = response.StrategyVariantEnabled
	o.ParitySymbol = response.ParitySymbol
	o.ExchangeName = response.ExchangeName

	out.TradeConfig = &o

	return
}

func (g *grpcServer) ListTradeConfig(ctx context.Context, in *pb.ListTradeConfigRequest) (out *pb.TradeConfigResponse, err error) {
	tradeConfigRepo := mysql.NewTradeConfigRepository(g.Db)
	useCases := usecase.NewListTradeConfigUsecase(tradeConfigRepo)

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

		out.TradeConfig = append(out.TradeConfig, &o)
	}

	return
}

func (g *grpcServer) CreateTradeConfig(ctx context.Context, in *pb.CreateTradeConfigRequest) (out *pb.TradeConfigResponse, err error) {
	tradeConfigRepo := mysql.NewTradeConfigRepository(g.Db)
	modalityRepo := mysql.NewModalityRepository(g.Db)
	exchangeRepo := mysql.NewExchangeRepository(g.Db)
	strategyRepo := mysql.NewStrategyRepository(g.Db)
	useCase := usecase.NewCreateTradeConfigUsecase(tradeConfigRepo, exchangeRepo, modalityRepo, strategyRepo)
	listUseCase := usecase.NewListTradeConfigUsecase(tradeConfigRepo)

	_, err = useCase.CreateTradeConfig(&usecasedto.InputTradeConfigDto{
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

	list, err := listUseCase.ListTradeConfig()
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
