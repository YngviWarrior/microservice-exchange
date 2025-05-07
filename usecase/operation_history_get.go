package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseGetOperationHistory struct {
	OperationHistoryRepo mysql.OperationHistoryRepositoryInterface
}

type UseCaseGetOperationHistoryInterface interface {
	GetOperationHistory(*usecasedto.InputGetOperationHistoryDto) (*usecasedto.OutputGetOperationHistoryDto, error)
}

func NewGetOperationHistoryUseCase(operationHistoryRepo mysql.OperationHistoryRepositoryInterface) UseCaseGetOperationHistoryInterface {
	return &usecaseGetOperationHistory{
		OperationHistoryRepo: operationHistoryRepo,
	}
}

func (u *usecaseGetOperationHistory) GetOperationHistory(in *usecasedto.InputGetOperationHistoryDto) (out *usecasedto.OutputGetOperationHistoryDto, err error) {
	operationHistory, err := u.OperationHistoryRepo.Get(in.OrderExchangeId)

	out = &usecasedto.OutputGetOperationHistoryDto{}

	out.OperationHistory = operationHistory.OperationHistory
	out.Operation = operationHistory.Operation
	out.TransactionType = operationHistory.TransactionType
	out.CoinPrice = operationHistory.CoinPrice
	out.CoinQuantity = operationHistory.CoinQuantity
	out.StablePrice = operationHistory.StablePrice
	out.StableQuantity = operationHistory.StableQuantity
	out.Fee = operationHistory.Fee
	out.OperationExchangeId = operationHistory.OperationExchangeId
	out.OperationExchangeStatus = operationHistory.OperationExchangeStatus

	return
}
