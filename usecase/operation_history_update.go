package usecase

import (
	"log"

	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseUpdateOperationHistory struct {
	OperationHistoryRepo mysql.OperationHistoryRepositoryInterface
}

type UseCaseUpdateOperationHistoryInterface interface {
	UpdateOperationHistory(*usecasedto.InputUpdateOperationHistoryDto) (*usecasedto.OutputUpdateOperationHistoryDto, error)
}

func NewUpdateOperationHistoryUseCase(operationHistoryRepo mysql.OperationHistoryRepositoryInterface) UseCaseUpdateOperationHistoryInterface {
	return &usecaseUpdateOperationHistory{
		OperationHistoryRepo: operationHistoryRepo,
	}
}

func (u *usecaseUpdateOperationHistory) UpdateOperationHistory(in *usecasedto.InputUpdateOperationHistoryDto) (out *usecasedto.OutputUpdateOperationHistoryDto, err error) {
	if !u.OperationHistoryRepo.Update(&repositorydto.InputOperationHistoryDto{
		OperationHistory:        in.OperationHistory,
		Operation:               in.Operation,
		TransactionType:         in.TransactionType,
		CoinPrice:               in.CoinPrice,
		CoinQuantity:            in.CoinQuantity,
		StablePrice:             in.StablePrice,
		StableQuantity:          in.StableQuantity,
		Fee:                     in.Fee,
		OperationExchangeId:     in.OperationExchangeId,
		OperationExchangeStatus: in.OperationExchangeStatus,
	}) {
		log.Println("uUOH 00: ", err)
		return
	}

	return
}
