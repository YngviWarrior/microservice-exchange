package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) GetAvgPriceByParityExchange(ctx context.Context, in *pb.GetAvgPriceByParityExchangeRequest) (out *pb.GetAvgPriceByParityExchangeResponse, err error) {
	avgRepo := mysql.NewAveragePriceRepository(g.Db)
	usecase := usecase.NewUsecaseGetAvgPriceByParityExchange(avgRepo)

	avg, err := usecase.GetAvgPriceByParityExchange(&usecasedto.InputGetAvgPriceByParityExchangeDto{
		Parity:   in.GetParity(),
		Exchange: in.GetExchange(),
	})
	if err != nil {
		return
	}

	out = &pb.GetAvgPriceByParityExchangeResponse{
		Parity:               avg.Parity,
		Exchange:             avg.Exchange,
		Day:                  avg.Day,
		DayRoc:               avg.DayRoc,
		DayUpdateTimestamp:   avg.DayUpdateTimestamp,
		Week:                 avg.Week,
		WeekRoc:              avg.WeekRoc,
		WeekUpdateTimestamp:  avg.WeekUpdateTimestamp,
		Month:                avg.Month,
		MonthRoc:             avg.MonthRoc,
		MonthUpdateTimestamp: avg.MonthUpdateTimestamp,
	}

	return
}

func (g *grpcServer) UpdateAveragePrice(ctx context.Context, in *pb.UpdateAveragePriceRequest) (out *pb.UpdateAveragePriceResponse, err error) {
	avgRepo := mysql.NewAveragePriceRepository(g.Db)
	usecase := usecase.NewUsecaseUpdateAveragePrice(avgRepo)

	_, err = usecase.UpdateAveragePrice(&usecasedto.InputUpdateAveragePriceDto{
		Parity:               in.GetParity(),
		Exchange:             in.GetExchange(),
		Day:                  in.GetDay(),
		DayRoc:               in.GetDayRoc(),
		DayUpdateTimestamp:   in.GetDayUpdateTimestamp(),
		Week:                 in.GetWeek(),
		WeekRoc:              in.GetWeekRoc(),
		WeekUpdateTimestamp:  in.GetWeekUpdateTimestamp(),
		Month:                in.GetMonth(),
		MonthRoc:             in.GetMonthRoc(),
		MonthUpdateTimestamp: in.GetMonthUpdateTimestamp(),
	})
	if err != nil {
		return
	}

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

	for _, avg := range avgList {
		var o pb.Candle

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
	}

	return
}
