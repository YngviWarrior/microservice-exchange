package grpcserver

import (
	"context"
	"fmt"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) GetAvgPriceByParityExchange(ctx context.Context, in *pb.GetAvgPriceByParityExchangeRequest) (out *pb.GetAvgPriceByParityExchangeResponse, err error) {
	avgRepo := mysql.NewAveragePriceRepository(g.Db)
	usecase := usecase.NewUsecaseGetAvgPriceByParityExchange(avgRepo)

	avg, err := usecase.GetAvgPriceByParityExchange(&usecasedto.InputGetAvgPriceByParityExchangeDto{})
	if err != nil {
		return
	}

	fmt.Println(avg)

	return
}

func (g *grpcServer) ListAvgPrices(ctx context.Context, in *pb.ListAvgPricesRequest) (out *pb.ListAvgPricesResponse, err error) {
	coinRepo := mysql.NewCoinRepository(g.Db)
	usecase := usecase.NewAvgPricesUseCase(coinRepo)

	avgList, err := usecase.ListAvgPrices(&usecasedto.InputAvgPricesDto{
		To:   in.To,
		From: in.From,
	})
	if err != nil {
		return
	}

	out = &pb.ListAvgPricesResponse{
		Candles: []*pb.Candle{},
	}

	var o pb.Candle
	for _, avg := range avgList {
		o.Parity = avg.Parity
		o.Exchange = avg.Exchange
		o.Mts = avg.Mts
		o.Open = avg.Open
		o.Close = avg.Close
		o.High = avg.High
		o.Low = avg.Low
		o.Volume = avg.Volume
		o.Roc = avg.Roc

		out.Candles = append(out.Candles, &o)
		fmt.Println(&o)
	}

	return
}
