package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) ListOperationByPeriod(ctx context.Context, in *pb.ListOperationByPeriodRequest) (out *pb.ListOperationByPeriodResponse, err error) {
	operationRepo := mysql.NewOperationRepository(g.Db)
	usecase := usecase.NewListOperationByPeriodUseCase(operationRepo)

	operations, err := usecase.ListOperationByPeriod(&usecasedto.InputListOperationByPeriodDto{
		MtsStart: in.GetMtsStart(),
		MtsEnd:   in.GetMtsEnd(),
	})
	if err != nil {
		return
	}

	out = &pb.ListOperationByPeriodResponse{
		Operations: []*pb.Operation{},
	}

	for _, v := range operations {
		var o pb.Operation

		o.Operation = uint64(v.Operation)
		o.User = v.User
		o.Parity = v.Parity
		o.Exchange = v.Exchange
		o.Strategy = uint64(v.Strategy)
		o.StrategyVariant = uint64(v.StrategyVariant)
		o.MtsStart = uint64(v.MtsStart)
		o.MtsFinish = uint64(v.MtsFinish)
		o.OpenPrice = v.OpenPrice
		o.ClosePrice = v.ClosePrice
		o.InvestedAmount = v.InvestedAmount
		o.ProfitAmount = v.ProfitAmount
		o.Profit = v.Profit
		o.Closed = v.Closed
		o.Audit = v.Audit
		o.Enabled = v.Enabled

		out.Operations = append(out.Operations, &o)
	}

	return
}

func (g *grpcServer) ListOperation(ctx context.Context, in *pb.ListOperationRequest) (out *pb.ListOperationResponse, err error) {
	operationRepo := mysql.NewOperationRepository(g.Db)
	usecase := usecase.NewListOperationUseCase(operationRepo)

	operations, err := usecase.ListOperation(&usecasedto.InputListOperationDto{
		UserId:          in.GetUser(),
		Strategy:        in.GetStrategy(),
		StrategyVariant: in.GetStrategyVariant(),
		Parity:          in.GetParity(),
		Exchange:        in.GetExchange(),
		Closed:          in.GetClosed(),
		Enabled:         in.GetEnabled(),
	})
	if err != nil {
		return
	}

	out = &pb.ListOperationResponse{
		Operations: []*pb.Operation{},
	}

	for _, v := range operations {
		var o pb.Operation

		o.Operation = uint64(v.Operation)
		o.User = v.User
		o.Parity = v.Parity
		o.Exchange = v.Exchange
		o.Strategy = uint64(v.Strategy)
		o.StrategyVariant = uint64(v.StrategyVariant)
		o.MtsStart = uint64(v.MtsStart)
		o.MtsFinish = uint64(v.MtsFinish)
		o.OpenPrice = v.OpenPrice
		o.ClosePrice = v.ClosePrice
		o.InvestedAmount = v.InvestedAmount
		o.ProfitAmount = v.ProfitAmount
		o.Profit = v.Profit
		o.Closed = v.Closed
		o.Audit = v.Audit
		o.Enabled = v.Enabled

		out.Operations = append(out.Operations, &o)
	}

	return
}
