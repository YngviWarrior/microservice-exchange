package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseCreateOperation struct {
	OperationRepo mysql.OperationRepositoryInterface
}

type UseCaseCreateOperationInterface interface {
	CreateOperation(*usecasedto.InputCreateOperationDto) (*usecasedto.OutputCreateOperationDto, error)
}

func NewCreateOperationUseCase(operationRepo mysql.OperationRepositoryInterface) UseCaseCreateOperationInterface {
	return &usecaseCreateOperation{
		OperationRepo: operationRepo,
	}
}

func (u *usecaseCreateOperation) CreateOperation(in *usecasedto.InputCreateOperationDto) (out *usecasedto.OutputCreateOperationDto, err error) {
	opId := u.OperationRepo.Create(&repositorydto.InputOperationDto{
		Operation:       in.Operation,
		User:            in.User,
		Parity:          in.Parity,
		Exchange:        in.Exchange,
		Strategy:        in.Strategy,
		StrategyVariant: in.StrategyVariant,
		MtsStart:        in.MtsStart,
		MtsFinish:       in.MtsFinish,
		OpenPrice:       in.OpenPrice,
		ClosePrice:      in.ClosePrice,
		InvestedAmount:  in.InvestedAmount,
		ProfitAmount:    in.ProfitAmount,
		Profit:          in.Profit,
		Closed:          in.Closed,
		Audit:           in.Audit,
		Enabled:         in.Enabled,
	})

	operation := u.OperationRepo.Get(uint64(opId))

	out = &usecasedto.OutputCreateOperationDto{
		Operation:       operation.Operation,
		User:            operation.User,
		Parity:          operation.Parity,
		Exchange:        operation.Exchange,
		Strategy:        operation.Strategy,
		StrategyVariant: operation.StrategyVariant,
		MtsStart:        operation.MtsStart,
		MtsFinish:       operation.MtsFinish,
		OpenPrice:       operation.OpenPrice,
		ClosePrice:      operation.ClosePrice,
		InvestedAmount:  operation.InvestedAmount,
		ProfitAmount:    operation.ProfitAmount,
		Profit:          operation.Profit,
		Closed:          operation.Closed,
		Audit:           operation.Audit,
		Enabled:         operation.Enabled,
	}

	return
}
