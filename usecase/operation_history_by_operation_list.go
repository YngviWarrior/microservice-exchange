package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListOperationHistoryByOperation struct {
	OperationHistoryRepo mysql.OperationHistoryRepositoryInterface
}

type UseCaseListOperationHistoryByOperationInterface interface {
	ListOperationHistoryByOperation(*usecasedto.InputListOperationHistoryByOperationDto) ([]*usecasedto.OutputListOperationHistoryByOperationDto, error)
}

func NewListOperationHistoryByOperationUseCase(operationHistoryRepo mysql.OperationHistoryRepositoryInterface) UseCaseListOperationHistoryByOperationInterface {
	return &usecaseListOperationHistoryByOperation{
		OperationHistoryRepo: operationHistoryRepo,
	}
}

func (u *usecaseListOperationHistoryByOperation) ListOperationHistoryByOperation(in *usecasedto.InputListOperationHistoryByOperationDto) (out []*usecasedto.OutputListOperationHistoryByOperationDto, err error) {
	list := u.OperationHistoryRepo.ListByOperation(in.Operation)

	out = []*usecasedto.OutputListOperationHistoryByOperationDto{}
	for _, v := range list {
		var o usecasedto.OutputListOperationHistoryByOperationDto

		o.OperationHistory = v.OperationHistory
		o.Operation = v.Operation
		o.TransactionType = v.TransactionType
		o.CoinPrice = v.CoinPrice
		o.CoinQuantity = v.CoinQuantity
		o.StablePrice = v.StablePrice
		o.StableQuantity = v.StableQuantity
		o.Fee = v.Fee
		o.OperationExchangeId = v.OperationExchangeId
		o.OperationExchangeStatus = v.OperationExchangeStatus

		out = append(out, &o)
	}

	return
}
