package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListOperationAll struct {
	OperationRepo mysql.OperationRepositoryInterface
}

type UseCaseListOperationAllInterface interface {
	ListOperationAll(*usecasedto.InputListOperationAllDto) ([]*usecasedto.OutputListOperationAllDto, error)
}

func NewListOperationAllUseCase(operationRepo mysql.OperationRepositoryInterface) UseCaseListOperationAllInterface {
	return &usecaseListOperationAll{
		OperationRepo: operationRepo,
	}
}

func (u *usecaseListOperationAll) ListOperationAll(in *usecasedto.InputListOperationAllDto) (out []*usecasedto.OutputListOperationAllDto, err error) {
	list := u.OperationRepo.ListAll()

	out = []*usecasedto.OutputListOperationAllDto{}
	for _, v := range list {
		var o usecasedto.OutputListOperationAllDto

		o.Operation = v.Operation
		o.User = v.User
		o.Parity = v.Parity
		o.Exchange = v.Exchange
		o.Strategy = v.Strategy
		o.StrategyName = v.StrategyName
		o.StrategyVariant = v.StrategyVariant
		o.StrategyVariantName = v.StrategyVariantName
		o.MtsStart = v.MtsStart
		o.MtsFinish = v.MtsFinish
		o.OpenPrice = v.OpenPrice
		o.ClosePrice = v.ClosePrice
		o.InvestedAmount = v.InvestedAmount
		o.ProfitAmount = v.ProfitAmount
		o.Profit = v.Profit
		o.Closed = v.Closed
		o.Audit = v.Audit
		o.Enabled = v.Enabled

		out = append(out, &o)
	}

	return
}
