package usecase

import (
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (u *usecase) ListExchanges() ([]*usecasedto.OutputListExchangeDto, error) {
	response := []*usecasedto.OutputListExchangeDto{}
	list := u.ExchangeRepo.List()

	for _, v := range list {
		x := &usecasedto.OutputListExchangeDto{}

		x.Exchange = v.Exchange
		x.Name = v.Name

		response = append(response, x)
	}

	return response, nil
}
