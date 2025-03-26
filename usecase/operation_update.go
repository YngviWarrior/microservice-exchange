package usecase

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseUpdateOperation struct {
	OperationRepo mysql.OperationRepositoryInterface
}

type UseCaseUpdateOperationInterface interface {
	UpdateOperation(*usecasedto.InputUpdateOperationDto) (*usecasedto.OutputUpdateOperationDto, error)
}

func NewUpdateOperationUseCase(operationRepo mysql.OperationRepositoryInterface) UseCaseUpdateOperationInterface {
	return &usecaseUpdateOperation{
		OperationRepo: operationRepo,
	}
}

func (u *usecaseUpdateOperation) UpdateOperation(in *usecasedto.InputUpdateOperationDto) (out *usecasedto.OutputUpdateOperationDto, err error) {
	if !u.OperationRepo.Update(&repositorydto.InputOperationDto{
		Operation:       uint64(in.Operation),
		User:            in.User,
		Parity:          in.Parity,
		Exchange:        in.Exchange,
		Strategy:        uint64(in.Strategy),
		StrategyVariant: uint64(in.StrategyVariant),
		MtsStart:        uint64(in.MtsStart),
		MtsFinish:       uint64(in.MtsFinish),
		OpenPrice:       in.OpenPrice,
		ClosePrice:      in.ClosePrice,
		InvestedAmount:  in.InvestedAmount,
		ProfitAmount:    in.ProfitAmount,
		Profit:          in.Profit,
		Closed:          in.Closed,
		Audit:           in.Audit,
		Enabled:         in.Enabled,
	}) {
		log.Println("uUO 00: ", err)
		return
	}

	return
}
