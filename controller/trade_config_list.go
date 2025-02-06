package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func (c *controller) ListTradeConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var send outputControllerDto

	output, err := c.UseCases.ListTradeConfig()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		send.Status = 0
		send.Errors = append(send.Errors, err.Error())
	} else {
		send.Status = 1
		send.Message = "Success"
		send.Data = output
	}

	jsonResp, err := json.Marshal(send)

	if err != nil {
		log.Panicf("CLTC 01: %s", err)
	}

	w.Write(jsonResp)
}
