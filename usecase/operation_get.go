package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseGetOperation struct {
	OperationRepo mysql.OperationRepositoryInterface
}

type UseCaseGetOperationInterface interface {
	GetOperation(*usecasedto.InputGetOperationDto) (*usecasedto.OutputGetOperationDto, error)
}

func NewGetOperationUseCase(operationRepo mysql.OperationRepositoryInterface) UseCaseGetOperationInterface {
	return &usecaseGetOperation{
		OperationRepo: operationRepo,
	}
}

func (u *usecaseGetOperation) GetOperation(in *usecasedto.InputGetOperationDto) (out *usecasedto.OutputGetOperationDto, err error) {
	operation := u.OperationRepo.Get(in.OperationId)

	out = &usecasedto.OutputGetOperationDto{}

	out.Operation = operation.Operation
	out.User = operation.User
	out.Parity = operation.Parity
	out.Exchange = operation.Exchange
	out.Strategy = operation.Strategy
	out.StrategyName = operation.StrategyName
	out.StrategyVariant = operation.StrategyVariant
	out.StrategyVariantName = operation.StrategyVariantName
	out.MtsStart = operation.MtsStart
	out.MtsFinish = operation.MtsFinish
	out.OpenPrice = operation.OpenPrice
	out.ClosePrice = operation.ClosePrice
	out.InvestedAmount = operation.InvestedAmount
	out.ProfitAmount = operation.ProfitAmount
	out.Profit = operation.Profit
	out.Closed = operation.Closed
	out.Audit = operation.Audit
	out.Enabled = operation.Enabled
	out.InTransaction = operation.InTransaction

	return
}
