package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListTransactionType struct {
	TransactionTypeRepo mysql.TransactionTypeRepositoryInterface
}

type UseCaseListTransactionTypeInterface interface {
	ListTransactionType(*usecasedto.InputListTransactionTypeDto) ([]*usecasedto.OutputListTransactionTypeDto, error)
}

func NewListTransactionTypeUseCase(transactionTypeRepo mysql.TransactionTypeRepositoryInterface) UseCaseListTransactionTypeInterface {
	return &usecaseListTransactionType{
		TransactionTypeRepo: transactionTypeRepo,
	}
}

func (u *usecaseListTransactionType) ListTransactionType(in *usecasedto.InputListTransactionTypeDto) (out []*usecasedto.OutputListTransactionTypeDto, err error) {
	list := u.TransactionTypeRepo.List(&repositorydto.InputTransactionTypeDto{})

	out = []*usecasedto.OutputListTransactionTypeDto{}
	for _, v := range list {
		var o usecasedto.OutputListTransactionTypeDto
		o.TransactionType = v.TransactionType
		o.Description = v.Description
		o.Active = v.Active

		out = append(out, &o)
	}

	return
}
