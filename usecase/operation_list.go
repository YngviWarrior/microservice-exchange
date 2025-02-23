package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListOperation struct {
	OperationRepo mysql.OperationRepositoryInterface
}

type UseCaseListOperationInterface interface {
	ListOperation(*usecasedto.InputListOperationDto) ([]*usecasedto.OutputListOperationDto, error)
}

func NewListOperationUseCase(operationRepo mysql.OperationRepositoryInterface) UseCaseListOperationInterface {
	return &usecaseListOperation{
		OperationRepo: operationRepo,
	}
}

func (u *usecaseListOperation) ListOperation(in *usecasedto.InputListOperationDto) (out []*usecasedto.OutputListOperationDto, err error) {
	list := u.OperationRepo.List(&repositorydto.InputOperationDto{
		User:            in.UserId,
		Strategy:        in.Strategy,
		StrategyVariant: in.StrategyVariant,
		Parity:          in.Parity,
		Exchange:        in.Exchange,
		Closed:          in.Closed,
		Enabled:         in.Enabled,
	})

	out = []*usecasedto.OutputListOperationDto{}
	for _, v := range list {
		var o usecasedto.OutputListOperationDto

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
