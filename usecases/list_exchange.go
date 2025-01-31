package usecases

import (
	"github.com/YngviWarrior/microservice-exchange/usecases/usecasedto"
)

func (u *usecases) ListExchanges() ([]*usecasedto.ListExchangeDto, error) {
	response := []*usecasedto.ListExchangeDto{}
	list := u.ExchangeRepo.List()

	for _, v := range list {
		x := &usecasedto.ListExchangeDto{}

		x.Exchange = v.Exchange
		x.Name = v.Name

		response = append(response, x)
	}

	return response, nil
}
