package controller

import (
	"net/http"

	"github.com/YngviWarrior/microservice-exchange/usecases"
)

type ControllerInterface interface {
	ListExchange(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	UseCases usecases.UseCasesInterface
}

type outputControllerDto struct {
	Status  int64    `json:"status,omitempty"`
	Message string   `json:"message,omitempty"`
	Data    any      `json:"data,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

func InitController(allUseCases usecases.UseCasesInterface) ControllerInterface {
	return &controller{
		UseCases: allUseCases,
	}
}
