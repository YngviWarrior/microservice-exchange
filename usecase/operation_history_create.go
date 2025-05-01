package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseCreateOperationHistory struct {
	OperationHistoryRepo mysql.OperationHistoryRepositoryInterface
}

type UseCaseCreateOperationHistoryInterface interface {
	CreateOperationHistory(*usecasedto.InputCreateOperationHistoryDto) (*usecasedto.OutputCreateOperationHistoryDto, error)
}

func NewCreateOperationHistoryUseCase(operationHistoryRepo mysql.OperationHistoryRepositoryInterface) UseCaseCreateOperationHistoryInterface {
	return &usecaseCreateOperationHistory{
		OperationHistoryRepo: operationHistoryRepo,
	}
}

func (u *usecaseCreateOperationHistory) CreateOperationHistory(in *usecasedto.InputCreateOperationHistoryDto) (out *usecasedto.OutputCreateOperationHistoryDto, err error) {
	if !u.OperationHistoryRepo.Create(&repositorydto.InputOperationHistoryDto{
		Operation:               in.Operation,
		TransactionType:         in.TransactionType,
		CoinPrice:               in.CoinPrice,
		CoinQuantity:            in.CoinQuantity,
		StablePrice:             in.StablePrice,
		StableQuantity:          in.StableQuantity,
		Fee:                     in.Fee,
		OperationExchangeId:     in.OperationExchangeId,
		OperationExchangeStatus: in.OperationExchangeStatus,
	}) {
		return nil, err
	}

	return
}
