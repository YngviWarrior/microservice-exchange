package controller

import (
	"net/http"

	"github.com/YngviWarrior/microservice-exchange/usecases"
)

type ControllerInterface interface {
	ListExchange(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	UseCases *usecases.UseCasesInterface
}

func InitController(allUseCases *usecases.UseCasesInterface) ControllerInterface {
	return &controller{
		UseCases: allUseCases,
	}
}
