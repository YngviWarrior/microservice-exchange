package internal

import (
	"fmt"
	"log"
	"strconv"
	"time"

	bybitstructs "github.com/YngviWarrior/bybit-sdk/byBitStructs"
	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
	"github.com/YngviWarrior/microservice-exchange/infra/internal/internaldto"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
)

type SellCoinParams struct {
	Operation    uint64
	ClosePrice   float64
	CoinQuantity float64
	Exchange     uint64
	OpAmount     float64
	OpFee        float64
	Symbol       string
	OrderType    string
}

func (b *strategy) SellCoin(params *internaldto.SellCoinParams) bool {
	switch params.Exchange {
	case 1: //Binance
	case 2: //ByBit
		p := &bybitstructs.OrderParams{
			Category:    "spot",
			Symbol:      params.Symbol,
			Side:        "Sell",
			OrderType:   params.OrderType,
			OrderPrice:  fmt.Sprintf("%.2f", params.ClosePrice),
			OrderQty:    fmt.Sprintf("%v", params.CoinQuantity),
			TimeInForce: "GTC",
		}

		res := b.Bybit.CreateOrder(p)
		log.Println("SELL ", res, " -> OP: ", params.Operation)

		if res.RetCode == 0 {
			orderId, _ := strconv.ParseInt(res.Result.OrderID, 10, 64)

			usecase.NewCreateOperationHistoryUseCase(b.OperationHistoryRepo).CreateOperationHistory(&usecasedto.InputCreateOperationHistoryDto{
				Operation:           params.Operation,
				TransactionType:     2,
				CoinPrice:           params.ClosePrice,
				CoinQuantity:        params.CoinQuantity,
				StablePrice:         params.OpAmount,
				StableQuantity:      params.OpAmount,
				Fee:                 params.OpFee,
				OperationExchangeId: uint64(orderId),
			})
		} else {
			discordService.NewDiscordWebhook().SendNotification(&discordstructs.Notification{
				ChannelUrl: "/1127721886214795325/h00kYmfUrFTAIIIJasJi2MOnD0jPRmja_NYXoC1gsxSv4pxKKFBQGWsr9T0FDGIf_RCk",
				Content:    fmt.Sprintf("(%v) <---> ERR Sell %v: Send %v <----> Get %v ", time.Now().Format("2006-01-02 15:04:05"), params.Operation, p, res),
			})
			return false
		}
	case 3: //TEST
		// var p bBR.OrderParams

		// p.Symbol = params.Symbol
		// p.Side = "SELL"
		// p.OrderType = "LIMIT"
		// p.OrderPrice = fmt.Sprintf("%v", params.ClosePrice)
		// p.OrderQty = fmt.Sprintf("%v", params.CoinQuantity)
		// p.TimeInForce = "GTC"

		// // res := bybittrade.SellOrder(&p)
		// // log.Println("SELL ", res)

		// // if res.RetCode == 0 {
		// orderId, _ := strconv.ParseInt("132546", 10, 64)
		// var pp repository.OperationHistoryDto

		// pp.Operation = params.Operation
		// pp.TransactionType = 1
		// pp.CoinPrice = params.ClosePrice
		// pp.CoinQuantity = p.OrderQty
		// pp.StablePrice = params.OpAmount
		// pp.StableQuantity = params.OpAmount
		// pp.Fee = 0
		// pp.OperationExchangeId = orderId

		// if !operationHistory.Create(tx, &pp) {
		// 	log.Panicln("SBC 02: ")
		// 	return false
		// }
		// // }
	}

	return true
}
