package usecases

import "github.com/YngviWarrior/microservice-exchange/usecases/usecasedto"

type UseCasesInterface interface {
	ListExchanges() *usecasedto.ListExchangeDto
}

type usecases struct{}

func InitUseCases() UseCasesInterface {
	return &usecases{}
}
