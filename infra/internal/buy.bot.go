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
	"github.com/YngviWarrior/microservice-exchange/utils"
)

func (i *strategy) BuyCoin(params *internaldto.BuyCoinParams) bool {
	switch params.Exchange {
	case 1:
	case 2:
		fmt.Println(params)
		res := i.Bybit.CreateOrder(&bybitstructs.OrderParams{
			Category:    "spot",
			Symbol:      params.Symbol,
			Side:        "Buy",
			OrderType:   params.OrderType,
			OrderPrice:  fmt.Sprintf("%v", params.ClosePrice),
			OrderQty:    fmt.Sprintf("%v", params.OpAmount),
			TimeInForce: "GTC",
		})

		log.Println("BUY ", res, " -> OP: ", params.Operation, params.Symbol)

		if res.RetCode == 0 {
			orderId, _ := strconv.ParseInt(res.Result.OrderID, 10, 64)
			opAmount, _ := strconv.ParseFloat(params.OpAmount, 64)
			fmt.Println("OpAmount:", opAmount)
			fmt.Println("Coin Price:", params.ClosePrice)
			fmt.Println("Fee:", params.OpFee)
			fmt.Println("Coin Quantity:", utils.Truncate((opAmount/params.ClosePrice)*(1-params.OpFee), 6))
			usecase.NewCreateOperationHistoryUseCase(i.OperationHistoryRepo).CreateOperationHistory(&usecasedto.InputCreateOperationHistoryDto{
				Operation:           params.Operation,
				TransactionType:     1,
				CoinPrice:           params.ClosePrice,
				CoinQuantity:        opAmount,
				StablePrice:         opAmount * params.ClosePrice,
				StableQuantity:      opAmount,
				Fee:                 params.OpFee,
				OperationExchangeId: uint64(orderId),
			})
		} else {
			discordService.NewDiscordWebhook().SendNotification(&discordstructs.Notification{
				ChannelUrl: "/1127721886214795325/h00kYmfUrFTAIIIJasJi2MOnD0jPRmja_NYXoC1gsxSv4pxKKFBQGWsr9T0FDGIf_RCk",
				Content:    fmt.Sprintf("(%v) <---> ERR Buy %v: Send %v <----> Get %v ", time.Now().Format("2006-01-02 15:04:05"), params.Operation, params.Operation, res),
			})
			return false
		}
	case 3:
		// var p bBR.OrderParams

		// p.Symbol = params.Symbol
		// p.Side = "BUY"
		// p.OrderType = "LIMIT"
		// p.OrderPrice = fmt.Sprintf("%v", params.ClosePrice)
		// p.OrderQty = fmt.Sprintf("%v", params.OpAmount)
		// p.TimeInForce = "GTC"

		// // res := bybittrade.BuyOrder(&p)

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
		// 	log.Panicln("BBC 02: ")
		// 	return false
		// }
		// }
	}

	return true
}
