package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseGetLastBuyRegisterByOperation struct {
	OperationHistoryRepo mysql.OperationHistoryRepositoryInterface
}

type UseCaseGetLastBuyRegisterByOperationInterface interface {
	GetLastBuyRegisterByOperation(*usecasedto.InputGetLastBuyRegisterByOperationDto) (*usecasedto.OutputGetLastBuyRegisterByOperationDto, error)
}

func NewGetLastBuyRegisterByOperationUseCase(operationHistoryRepo mysql.OperationHistoryRepositoryInterface) UseCaseGetLastBuyRegisterByOperationInterface {
	return &usecaseGetLastBuyRegisterByOperation{
		OperationHistoryRepo: operationHistoryRepo,
	}
}

func (u *usecaseGetLastBuyRegisterByOperation) GetLastBuyRegisterByOperation(in *usecasedto.InputGetLastBuyRegisterByOperationDto) (out *usecasedto.OutputGetLastBuyRegisterByOperationDto, err error) {
	coinQuantity, fee, status := u.OperationHistoryRepo.GetLastBuyRegisterByOperation(in.Operation)

	out = &usecasedto.OutputGetLastBuyRegisterByOperationDto{}
	out.CoinQuantity = coinQuantity
	out.Fee = fee
	out.Status = status

	return
}
