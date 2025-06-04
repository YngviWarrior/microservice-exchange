package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListOperationEnabled struct {
	OperationRepo mysql.OperationRepositoryInterface
}

type UseCaseListOperationEnabledInterface interface {
	ListOperationEnabled(*usecasedto.InputListOperationEnabledDto) ([]*usecasedto.OutputListOperationEnabledDto, error)
}

func NewListOperationEnabledUseCase(operationRepo mysql.OperationRepositoryInterface) UseCaseListOperationEnabledInterface {
	return &usecaseListOperationEnabled{
		OperationRepo: operationRepo,
	}
}

func (u *usecaseListOperationEnabled) ListOperationEnabled(in *usecasedto.InputListOperationEnabledDto) (out []*usecasedto.OutputListOperationEnabledDto, err error) {
	operations := u.OperationRepo.ListEnabled()

	out = []*usecasedto.OutputListOperationEnabledDto{}
	for _, operation := range operations {
		o := &usecasedto.OutputListOperationEnabledDto{}

		o.Operation = operation.Operation
		o.User = operation.User
		o.Parity = operation.Parity
		o.ParitySymbol = operation.ParitySymbol
		o.Exchange = operation.Exchange
		o.ExchangeName = operation.ExchangeName
		o.Strategy = operation.Strategy
		o.StrategyName = operation.StrategyName
		o.StrategyVariant = operation.StrategyVariant
		o.StrategyVariantName = operation.StrategyVariantName
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
		o.InTransaction = operation.InTransaction

		out = append(out, o)
	}

	return
}
