package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) ListAllOperation(ctx context.Context, in *pb.ListAllOperationRequest) (out *pb.ListAllOperationResponse, err error) {
	operationRepo := mysql.NewOperationRepository(g.Db)
	usecase := usecase.NewListOperationAllUseCase(operationRepo)

	operations, err := usecase.ListOperationAll(&usecasedto.InputListOperationAllDto{})
	if err != nil {
		return
	}

	out = &pb.ListAllOperationResponse{
		Operations: []*pb.OperationJoin{},
	}

	for _, v := range operations {
		var o pb.OperationJoin

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

func (g *grpcServer) GetOperation(ctx context.Context, in *pb.GetOperationRequest) (out *pb.GetOperationResponse, err error) {
	operationRepo := mysql.NewOperationRepository(g.Db)
	usecase := usecase.NewGetOperationUseCase(operationRepo)

	operation, err := usecase.GetOperation(&usecasedto.InputGetOperationDto{
		OperationId: in.GetOperationId(),
	})
	if err != nil {
		return
	}

	out = &pb.GetOperationResponse{
		Operations: &pb.Operation{},
	}

	var o pb.Operation

	o.Operation = uint64(operation.Operation)
	o.User = operation.User
	o.Parity = operation.Parity
	o.Exchange = operation.Exchange
	o.Strategy = uint64(operation.Strategy)
	o.StrategyVariant = uint64(operation.StrategyVariant)
	o.MtsStart = uint64(operation.MtsStart)
	o.MtsFinish = uint64(operation.MtsFinish)
	o.OpenPrice = operation.OpenPrice
	o.ClosePrice = operation.ClosePrice
	o.InvestedAmount = operation.InvestedAmount
	o.ProfitAmount = operation.ProfitAmount
	o.Profit = operation.Profit
	o.Closed = operation.Closed
	o.Audit = operation.Audit
	o.Enabled = operation.Enabled

	out.Operations = &o

	return
}

func (g *grpcServer) UpdateOperation(ctx context.Context, in *pb.UpdateOperationRequest) (out *pb.UpdateOperationResponse, err error) {
	operationRepo := mysql.NewOperationRepository(g.Db)
	usecase := usecase.NewUpdateOperationUseCase(operationRepo)

	op := in.GetOperation()

	operation, err := usecase.UpdateOperation(&usecasedto.InputUpdateOperationDto{
		Operation:       op.GetOperation(),
		User:            op.GetUser(),
		Parity:          op.GetParity(),
		Exchange:        op.GetExchange(),
		Strategy:        op.GetStrategy(),
		StrategyVariant: op.GetStrategyVariant(),
		MtsStart:        op.GetMtsStart(),
		MtsFinish:       op.GetMtsFinish(),
		OpenPrice:       op.GetOpenPrice(),
		ClosePrice:      op.GetClosePrice(),
		InvestedAmount:  op.GetInvestedAmount(),
		ProfitAmount:    op.GetProfitAmount(),
		Profit:          op.GetProfit(),
		Closed:          op.GetClosed(),
		Audit:           op.GetAudit(),
		Enabled:         op.GetEnabled(),
	})
	if err != nil {
		return
	}

	out = &pb.UpdateOperationResponse{
		Operation: &pb.Operation{},
	}

	var o pb.Operation

	o.Operation = uint64(operation.Operation)
	o.User = operation.User
	o.Parity = operation.Parity
	o.Exchange = operation.Exchange
	o.Strategy = uint64(operation.Strategy)
	o.StrategyVariant = uint64(operation.StrategyVariant)
	o.MtsStart = uint64(operation.MtsStart)
	o.MtsFinish = uint64(operation.MtsFinish)
	o.OpenPrice = operation.OpenPrice
	o.ClosePrice = operation.ClosePrice
	o.InvestedAmount = operation.InvestedAmount
	o.ProfitAmount = operation.ProfitAmount
	o.Profit = operation.Profit
	o.Closed = operation.Closed
	o.Audit = operation.Audit
	o.Enabled = operation.Enabled

	out.Operation = &o

	return
}

func (g *grpcServer) CreateOperation(ctx context.Context, in *pb.CreateOperationRequest) (out *pb.CreateOperationResponse, err error) {
	operationRepo := mysql.NewOperationRepository(g.Db)
	usecase := usecase.NewCreateOperationUseCase(operationRepo)

	op := in.GetOperation()

	operation, err := usecase.CreateOperation(&usecasedto.InputCreateOperationDto{
		Operation:       op.GetOperation(),
		User:            op.GetUser(),
		Parity:          op.GetParity(),
		Exchange:        op.GetExchange(),
		Strategy:        op.GetStrategy(),
		StrategyVariant: op.GetStrategyVariant(),
		MtsStart:        op.GetMtsStart(),
		MtsFinish:       op.GetMtsFinish(),
		OpenPrice:       op.GetOpenPrice(),
		ClosePrice:      op.GetClosePrice(),
		InvestedAmount:  op.GetInvestedAmount(),
		ProfitAmount:    op.GetProfitAmount(),
		Profit:          op.GetProfit(),
		Closed:          op.GetClosed(),
		Audit:           op.GetAudit(),
		Enabled:         op.GetEnabled(),
	})
	if err != nil {
		return
	}

	out = &pb.CreateOperationResponse{
		Operation: &pb.Operation{},
	}

	var o pb.Operation

	o.Operation = operation.Operation
	o.User = operation.User
	o.Parity = operation.Parity
	o.Exchange = operation.Exchange
	o.Strategy = operation.Strategy
	o.StrategyVariant = operation.StrategyVariant
	o.MtsStart = operation.MtsStart
	o.MtsFinish = operation.MtsFinish
	o.OpenPrice = operation.OpenPrice
	o.ClosePrice = operation.ClosePrice
	o.InvestedAmount = operation.InvestedAmount
	o.ProfitAmount = operation.ProfitAmount
	o.Profit = operation.Profit
	o.Closed = operation.Closed
	o.Audit = operation.Audit
	o.Enabled = operation.Enabled

	out.Operation = &o

	return
}
