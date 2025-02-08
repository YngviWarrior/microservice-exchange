package controller

import (
	"encoding/json"
	"log"
	"net/http"

	controllerdto "github.com/YngviWarrior/microservice-exchange/controller/controller_dto"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

func (c *controller) CreateTradeConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input controllerdto.InputTradeConfigDto
	var send outputControllerDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Panicf("CCTC 00: %s", err)
		w.WriteHeader(http.StatusInternalServerError)

		send.Errors = append(send.Errors, "invalid primitive type")
		jsonResp, err := json.Marshal(send)

		if err != nil {
			log.Fatalf("CCTC00.1: %s", err)
		}

		w.Write(jsonResp)
		return
	}

	if !c.InputValidation(w, input) {
		return
	}

	output, err := c.UseCases.CreateTradeConfig(&usecasedto.InputTradeConfigDto{
		Modality:                *input.Modality,
		Strategy:                *input.Strategy,
		StrategyVariant:         *input.StrategyVariant,
		Parity:                  *input.Parity,
		Exchange:                *input.Exchange,
		OperationQuantity:       *input.OperationQuantity,
		OperationAmount:         *input.OperationAmount,
		DefaultProfitPercentage: *input.DefaultProfitPercentage,
		WalletValueLimit:        *input.WalletValueLimit,
		Enabled:                 *input.Enabled,
	})

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
		log.Panicf("CLE 01: %s", err)
	}

	w.Write(jsonResp)
}
