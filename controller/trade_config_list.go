package controller

import (
	"net/http"

	controllerdto "github.com/YngviWarrior/microservice-exchange/controller/controller_dto"
)

func (c *controller) ListTradeConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var send outputControllerDto

	output, err := c.UseCases.ListTradeConfig()

	var out []*controllerdto.OutputTradeConfigDto
	for _, v := range output {
		out = append(out, &controllerdto.OutputTradeConfigDto{
			TradeConfig:             v.TradeConfig,
			Modality:                v.Modality,
			Strategy:                v.Strategy,
			StrategyVariant:         v.StrategyVariant,
			Parity:                  v.Parity,
			Exchange:                v.Exchange,
			OperationQuantity:       v.OperationQuantity,
			OperationAmount:         v.OperationAmount,
			DefaultProfitPercentage: v.DefaultProfitPercentage,
			WalletValueLimit:        v.WalletValueLimit,
			Enabled:                 v.Enabled,
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
