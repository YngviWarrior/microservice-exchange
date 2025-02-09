package controller

import (
	"net/http"

	controllerdto "github.com/YngviWarrior/microservice-exchange/controller/controller_dto"
)

func (c *controller) ListExchange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var send outputControllerDto

	output, err := c.UseCases.ListExchanges()

	var out []controllerdto.OutputListExchange
	for _, v := range output {
		out = append(out, controllerdto.OutputListExchange{
			Exchange: v.Exchange,
			Name:     v.Name,
		})
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		send.Status = 0
		send.Errors = append(send.Errors, err.Error())
	} else {
		send.Status = 1
		send.Message = "Success"
		send.Data = out
	}

	c.FormatResponse(w, send)
}
