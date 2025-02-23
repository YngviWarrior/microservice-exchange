package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql/repositorydto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type usecaseListParity struct {
	ParityRepo mysql.ParityRepositoryInterface
}

type UsecaseListParityInteface interface {
	ListParity(in *usecasedto.InputListParityDto) (out []*usecasedto.OutputListParityDto, err error)
}

func NewUsecaseListParity(parityRepo mysql.ParityRepositoryInterface) UsecaseListParityInteface {
	return &usecaseListParity{
		ParityRepo: parityRepo,
	}
}

func (u *usecaseListParity) ListParity(in *usecasedto.InputListParityDto) (out []*usecasedto.OutputListParityDto, err error) {
	parities := u.ParityRepo.List(&repositorydto.InputParityDto{})

	for _, v := range parities {
		var o usecasedto.OutputListParityDto

		o.Active = v.Active
		o.Parity = v.Parity
		o.Symbol = v.Symbol

		out = append(out, &o)
	}

	return
}
