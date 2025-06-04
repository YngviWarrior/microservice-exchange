package grpcserver

import (
	"context"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/grpcServer/proto/pb"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (g *grpcServer) ListOperationHistoryByOperation(ctx context.Context, in *pb.ListOperationHistoryByOperationRequest) (out *pb.ListOperationHistoryByOperationResponse, err error) {
	operationHistoryRepo := mysql.NewOperationHistoryRepository(g.Db)
	usecase := usecase.NewListOperationHistoryByOperationUseCase(operationHistoryRepo)

	response, err := usecase.ListOperationHistoryByOperation(&usecasedto.InputListOperationHistoryByOperationDto{
		Operation: in.GetOperation(),
	})
	if err != nil {
		return
	}

	out = &pb.ListOperationHistoryByOperationResponse{
		OperationHistory: []*pb.OperationHistory{},
	}

	for _, v := range response {
		out.OperationHistory = append(out.OperationHistory, &pb.OperationHistory{
			OperationHistory:        v.OperationHistory,
			Operation:               v.Operation,
			TransactionType:         v.TransactionType,
			CoinPrice:               v.CoinPrice,
			CoinQuantity:            v.CoinQuantity,
			StablePrice:             v.StablePrice,
			StableQuantity:          v.StableQuantity,
			Fee:                     v.Fee,
			OperationExchangeId:     v.OperationExchangeId,
			OperationExchangeStatus: v.OperationExchangeStatus,
		})
	}

	return
}

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
		Status:       response.Status,
	}

	return
}

func (g *grpcServer) GetOperationHistory(ctx context.Context, in *pb.GetOperationHistoryRequest) (out *pb.GetOperationHistoryResponse, err error) {
	operationHistoryRepo := mysql.NewOperationHistoryRepository(g.Db)
	usecase := usecase.NewGetOperationHistoryUseCase(operationHistoryRepo)

	response, err := usecase.GetOperationHistory(&usecasedto.InputGetOperationHistoryDto{
		OrderExchangeId: in.GetOrderId(),
	})
	if err != nil {
		return
	}

	out = &pb.GetOperationHistoryResponse{
		OperationHistory: &pb.OperationHistory{
			OperationHistory:        response.OperationHistory,
			Operation:               response.Operation,
			TransactionType:         response.TransactionType,
			CoinPrice:               response.CoinPrice,
			CoinQuantity:            response.CoinQuantity,
			StablePrice:             response.StablePrice,
			StableQuantity:          response.StableQuantity,
			Fee:                     response.Fee,
			OperationExchangeId:     response.OperationExchangeId,
			OperationExchangeStatus: response.OperationExchangeStatus,
		},
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

func (g *grpcServer) UpdateOperationHistory(ctx context.Context, in *pb.UpdateOperationHistoryRequest) (out *pb.UpdateOperationHistoryResponse, err error) {
	operationHistoryRepo := mysql.NewOperationHistoryRepository(g.Db)
	usecase := usecase.NewUpdateOperationHistoryUseCase(operationHistoryRepo)

	_, err = usecase.UpdateOperationHistory(&usecasedto.InputUpdateOperationHistoryDto{
		OperationHistory:        in.GetOperationHistory().GetOperationHistory(),
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
