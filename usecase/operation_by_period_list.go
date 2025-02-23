package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListOperationByPeriod struct {
	OperationRepo mysql.OperationRepositoryInterface
}

type UseCaseListOperationByPeriodInterface interface {
	ListOperationByPeriod(*usecasedto.InputListOperationByPeriodDto) ([]*usecasedto.OutputListOperationByPeriodDto, error)
}

func NewListOperationByPeriodUseCase(operationRepo mysql.OperationRepositoryInterface) UseCaseListOperationByPeriodInterface {
	return &usecaseListOperationByPeriod{
		OperationRepo: operationRepo,
	}
}

func (u *usecaseListOperationByPeriod) ListOperationByPeriod(in *usecasedto.InputListOperationByPeriodDto) (out []*usecasedto.OutputListOperationByPeriodDto, err error) {
	list := u.OperationRepo.ListByPeriod(int64(in.MtsStart), int64(in.MtsEnd))

	out = []*usecasedto.OutputListOperationByPeriodDto{}
	for _, v := range list {
		var o usecasedto.OutputListOperationByPeriodDto

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
