package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) ListCandleLimit(ctx context.Context, in *pb.ListCandleLimitRequest) (out *pb.ListCandleLimitResponse, err error) {
	candleRepo := mysql.NewCandleRepository(g.Db)
	usecase := usecase.NewListCandleLimit(candleRepo)

	candles, err := usecase.ListCandleLimit(&usecasedto.InputListCandleLimitDto{
		Parity:   in.GetParity(),
		Exchange: in.GetExchange(),
		Limit:    in.GetLimit(),
	})

	out = &pb.ListCandleLimitResponse{
		Candles: []*pb.Candle{},
	}

	for _, candle := range candles.Candles {
		var c pb.Candle

		c.Parity = candle.Parity
		c.Exchange = candle.Exchange
		c.Mts = candle.Mts
		c.Open = candle.Open
		c.Close = candle.Close
		c.High = candle.High
		c.Low = candle.Low
		c.Volume = candle.Volume
		c.Roc = candle.Roc

		out.Candles = append(out.Candles, &c)
	}

	return
}

func (g *grpcServer) CreateCandles(ctx context.Context, in *pb.CreateCandlesRequest) (out *pb.CreateCandlesResponse, err error) {
	candleRepo := mysql.NewCandleRepository(g.Db)
	usecase := usecase.NewCreateCandles(candleRepo)

	param := &usecasedto.InputCreateCandlesDto{}
	for _, candle := range in.GetCandles() {
		var c usecasedto.Candles

		c.Parity = candle.GetParity()
		c.Exchange = candle.GetExchange()
		c.Mts = candle.GetMts()
		c.Open = candle.GetOpen()
		c.Close = candle.GetClose()
		c.High = candle.GetHigh()
		c.Low = candle.GetLow()
		c.Volume = candle.GetVolume()

		param.Candles = append(param.Candles, c)
	}

	usecase.CreateCandles(param)
	return
}

func (g *grpcServer) GetCandleFirstMts(ctx context.Context, in *pb.GetCandleFirstMtsRequest) (out *pb.GetCandleFirstMtsResponse, err error) {
	candleRepo := mysql.NewCandleRepository(g.Db)
	usecase := usecase.NewGetCandleFirstMts(candleRepo)

	candle, err := usecase.GetCandleFirstMts(&usecasedto.InputGetCandleFirstMtsDto{
		Parity:   in.GetParity(),
		Exchange: in.GetExchange(),
	})
	if err != nil {
		return
	}

	out = &pb.GetCandleFirstMtsResponse{
		Candles: []*pb.Candle{},
	}

	var o pb.Candle

	o.Parity = candle.Parity
	o.Exchange = candle.Exchange
	o.Mts = candle.Mts
	o.Open = candle.Open
	o.Close = candle.Close
	o.High = candle.High
	o.Low = candle.Low
	o.Volume = candle.Volume
	o.Roc = candle.Roc

	out.Candles = append(out.Candles, &o)

	return
}
