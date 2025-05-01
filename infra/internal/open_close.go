package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	discordService "github.com/YngviWarrior/discord-webhook"
	discordstructs "github.com/YngviWarrior/discord-webhook/discordStructs"
	"github.com/YngviWarrior/microservice-exchange/infra/internal/internaldto"
	"github.com/YngviWarrior/microservice-exchange/usecase"
	"github.com/YngviWarrior/microservice-exchange/usecase/usecasedto"
	"github.com/YngviWarrior/microservice-exchange/utils"
)

func (u *strategy) OpenClose(in *internaldto.InputStrategyOpenCloseDto) (out []*internaldto.OutputStrategyOpenCloseDto, err error) {
	var newRegister bool
	var stableCoin uint64
	switch in.TradeConfig.Parity {
	case 1, 4:
		stableCoin = 1
	case 3:
		stableCoin = 2
	}

	wallet, err := usecase.NewGetWalletWithCoinUsecase(u.WalletRepo).GetWalletWithCoin(&usecasedto.InputWalletWithCoinDto{
		Exchange: in.TradeConfig.Exchange,
		Coin:     stableCoin,
	})
	if err != nil {
		log.Panicln("OC 00: ", err)
	}
	fmt.Println(wallet)

	operations, err := usecase.NewListOperationUseCase(u.OperationRepo).ListOperation(&usecasedto.InputListOperationDto{
		UserId:          in.TradeConfig.User,
		Strategy:        in.TradeConfig.Strategy,
		StrategyVariant: in.TradeConfig.StrategyVariant,
		Parity:          in.TradeConfig.Parity,
		Exchange:        in.TradeConfig.Exchange,
		Closed:          false,
		Enabled:         true,
	})
	if err != nil {
		log.Panicln("OC 01: ", err)
	}
	fmt.Println(operations)

	// fmt.Println("Open Operations: ", len(operations.Operations)
	if len(operations) < int(in.TradeConfig.OperationQuantity) && wallet.Amount > (in.TradeConfig.OperationAmount+in.TradeConfig.WalletValueLimit) {
		op, err := usecase.NewCreateOperationUseCase(u.OperationRepo).CreateOperation(&usecasedto.InputCreateOperationDto{
			User:            in.TradeConfig.User,
			Parity:          in.TradeConfig.Parity,
			Exchange:        in.TradeConfig.Exchange,
			Strategy:        in.TradeConfig.Strategy,
			StrategyVariant: in.TradeConfig.StrategyVariant,
			InvestedAmount:  in.TradeConfig.OperationAmount,
			Enabled:         true,
		})
		if err != nil {
			log.Panicln("OC 03: ", err)
		}

		if op.Operation == 0 {
			log.Panicln("OC 04: ")
		}

		discordService.NewDiscordWebhook().SendNotification(
			&discordstructs.Notification{
				ChannelUrl: "/1127722725318856714/31CT1E3rKDE-U1Ip-4nuYX_aIKJKGpKu9D3_p122FRMD5zbKs719t4YFdlV1LtTNDfHL",
				Content:    fmt.Sprintf("(%v) <---> Registering New Order! %v Strategy %v (%v)", time.Now().Format("2006-01-02 15:04:05"), in.TradeConfig.ParitySymbol, in.TradeConfig.StrategyName, in.TradeConfig.StrategyVariantName),
			},
		)

		newRegister = true
	}

	for _, operation := range operations {
		lastRegister, err := usecase.NewGetLastBuyRegisterByOperationUseCase(u.OperationHistoryRepo).GetLastBuyRegisterByOperation(&usecasedto.InputGetLastBuyRegisterByOperationDto{
			Operation: uint64(operation.Operation),
		})
		if err != nil {
			log.Panicln("OC 05: ", err)
		}
		fmt.Println("lastRegister: ", lastRegister)

		sellCotation := utils.FindSellCotation(operation.InvestedAmount, lastRegister.CoinQuantity, lastRegister.Fee, operation.OpenPrice, in.TradeConfig.DefaultProfitPercentage)

		currentPrice := in.ClosePrice
		var isProfitable bool = (currentPrice >= sellCotation) && sellCotation != 0
		// fmt.Printf("%v OP: %v -> SellCotation %v \n", time.Now().Format("2006-01-02 15:04:05"), operation.Operation, sellCotation)
		if wallet.Amount > (in.TradeConfig.OperationAmount+in.TradeConfig.WalletValueLimit) && lastRegister.CoinQuantity == 0 {
			var bc internaldto.BuyCoinParams

			bc.Operation = operation.Operation
			bc.ClosePrice = currentPrice
			bc.Exchange = in.TradeConfig.Exchange
			bc.Symbol = in.TradeConfig.ParitySymbol
			bc.OrderType = "Limit"
			bc.OpFee = 0.001

			switch operation.Parity {
			case 1: // BTCUSDT
				bc.OpAmount = fmt.Sprintf("%.6f", utils.Truncate(in.TradeConfig.OperationAmount/currentPrice, 5))
			case 4: // ETHUSDT
				bc.OpAmount = fmt.Sprintf("%.5f", utils.Truncate(in.TradeConfig.OperationAmount/currentPrice, 4))
			}

			switch os.Getenv("ENVIROMENT") {
			case "server", "local", "testnet":
				if u.BuyCoin(&bc) {
					usecase.NewUpdateWalletUseCase(u.WalletRepo).UpdateWallet([]*usecasedto.InputUpdateWalletDto{
						{
							Wallet:   wallet.Wallet,
							Exchange: wallet.Exchange,
							Coin:     wallet.Coin,
							Amount:   wallet.Amount + (-1 * in.TradeConfig.OperationAmount),
						},
					})

					usecase.NewUpdateOperationUseCase(u.OperationRepo).UpdateOperation(&usecasedto.InputUpdateOperationDto{
						Operation:       operation.Operation,
						User:            operation.User,
						Parity:          operation.Parity,
						Exchange:        operation.Exchange,
						Strategy:        operation.Strategy,
						StrategyVariant: operation.StrategyVariant,
						MtsStart:        operation.MtsStart,
						InvestedAmount:  operation.InvestedAmount,
						OpenPrice:       currentPrice,
						ClosePrice:      operation.ClosePrice,
						Profit:          in.TradeConfig.DefaultProfitPercentage,
						ProfitAmount:    operation.ProfitAmount,
						Closed:          false,
						Audit:           operation.Audit,
						Enabled:         true,
					})

					newRegister = true
				}
			}
		} else if isProfitable && lastRegister.CoinQuantity != 0 {
			var p internaldto.SellCoinParams

			p.Operation = operation.Operation
			p.ClosePrice = currentPrice
			p.Exchange = in.TradeConfig.Exchange
			p.OpAmount = operation.InvestedAmount
			p.Symbol = in.TradeConfig.ParitySymbol
			p.OrderType = "Limit"
			p.OpFee = 0.001

			switch operation.Parity {
			case 1: // BTCUSDT
				p.CoinQuantity = utils.Truncate(lastRegister.CoinQuantity, 5)
			case 4: // ETHUSDT
				p.CoinQuantity = utils.Truncate(lastRegister.CoinQuantity, 4)
			}

			switch os.Getenv("ENVIROMENT") {
			case "server", "local", "testnet":
				if u.SellCoin(&p) {
					usecase.NewUpdateWalletUseCase(u.WalletRepo).UpdateWallet([]*usecasedto.InputUpdateWalletDto{
						{
							Wallet:   wallet.Wallet,
							Exchange: wallet.Exchange,
							Coin:     wallet.Coin,
							Amount:   (operation.InvestedAmount * (1 + in.TradeConfig.DefaultProfitPercentage)),
						},
					})

					usecase.NewUpdateOperationUseCase(
						u.OperationRepo,
					).UpdateOperation(&usecasedto.InputUpdateOperationDto{
						Operation:       operation.Operation,
						User:            operation.User,
						Parity:          operation.Parity,
						Exchange:        operation.Exchange,
						Strategy:        operation.Strategy,
						StrategyVariant: operation.StrategyVariant,
						MtsStart:        operation.MtsStart,
						MtsFinish:       uint64(time.Now().UnixMicro()),
						InvestedAmount:  operation.InvestedAmount * (1 + in.TradeConfig.DefaultProfitPercentage),
						OpenPrice:       currentPrice,
						ClosePrice:      currentPrice,
						Profit:          in.TradeConfig.DefaultProfitPercentage,
						ProfitAmount:    (currentPrice * p.CoinQuantity) - operation.InvestedAmount,
						Closed:          false,
						Audit:           operation.Audit,
						Enabled:         true,
					})

					newRegister = true
				}
			}
		}
	}

	if newRegister {
		time.Sleep(time.Millisecond * 500)
	}

	return
}
