package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) GetLastBuyRegisterByOperation(ctx context.Context, in *pb.GetLastBuyRegisterByOperationRequest) (out *pb.GetLastBuyRegisterByOperationResponse, err error) {
	operationHistoryRepo := mysql.NewOperationHistoryRepository(g.Db)
	usecase := usecase.NewGetLastBuyRegisterByOperationUseCase(operationHistoryRepo)

	response, err := usecase.GetLastBuyRegisterByOperation(&usecasedto.InputGetLastBuyRegisterByOperationDto{
		Operation: in.GetOperation(),
	})
	if err != nil {
		return
	}

	out = &pb.GetLastBuyRegisterByOperationResponse{
		CoinQuantity: response.CoinQuantity,
		Fee:          response.Fee,
	}

	return
}
