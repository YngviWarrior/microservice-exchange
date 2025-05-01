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

func (g *grpcServer) CreateOperationHistory(ctx context.Context, in *pb.CreateOperationHistoryRequest) (out *pb.CreateOperationHistoryResponse, err error) {
	operationHistoryRepo := mysql.NewOperationHistoryRepository(g.Db)
	usecase := usecase.NewCreateOperationHistoryUseCase(operationHistoryRepo)

	_, err = usecase.CreateOperationHistory(&usecasedto.InputCreateOperationHistoryDto{
		Operation:               in.GetOperationHistory().GetOperation(),
		TransactionType:         in.GetOperationHistory().GetTransactionType(),
		CoinPrice:               in.GetOperationHistory().GetCoinPrice(),
		CoinQuantity:            in.GetOperationHistory().GetCoinQuantity(),
		StablePrice:             in.GetOperationHistory().GetStablePrice(),
		StableQuantity:          in.GetOperationHistory().GetStableQuantity(),
		Fee:                     in.GetOperationHistory().GetFee(),
		OperationExchangeId:     in.GetOperationHistory().GetOperationExchangeId(),
		OperationExchangeStatus: in.GetOperationHistory().GetOperationExchangeStatus(),
	})
	if err != nil {
		return
	}

	return
}
