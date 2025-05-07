package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) ListTransactionType(ctx context.Context, in *pb.ListTransactionTypeRequest) (out *pb.ListTransactionTypeResponse, err error) {
	transactionTypeRepo := mysql.NewTransactionTypeRepository(g.Db)
	usecase := usecase.NewListTransactionTypeUseCase(transactionTypeRepo)

	list, err := usecase.ListTransactionType(&usecasedto.InputListTransactionTypeDto{})
	if err != nil {
		return
	}

	out = &pb.ListTransactionTypeResponse{
		TransactionTypes: []*pb.TransactionType{},
	}

	for _, v := range list {
		var o pb.TransactionType

		o.Active = v.Active
		o.Description = v.Description
		o.Active = v.Active

		out.TransactionTypes = append(out.TransactionTypes, &o)
	}

	return
}
