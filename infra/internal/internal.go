package internal

import (
	"os"

	bybitSDK "github.com/YngviWarrior/bybit-sdk"
	"github.com/YngviWarrior/microservice-exchange/infra/database/mysql"
	"github.com/YngviWarrior/microservice-exchange/infra/internal/internaldto"
)

type strategy struct {
	Bybit bybitSDK.BybitServiceInterface

	OperationRepo        mysql.OperationRepositoryInterface
	OperationHistoryRepo mysql.OperationHistoryRepositoryInterface
	WalletRepo           mysql.WalletRepositoryInterface
	ParityRepo           mysql.ParityRepositoryInterface
	AverageRepo          mysql.AveragePriceRepositoryInterface
}

type StrategyInteface interface {
	BuyCoin(in *internaldto.BuyCoinParams) bool
	SellCoin(in *internaldto.SellCoinParams) bool
	AveragePrice(in *internaldto.InputStrategyAveragePriceDto) (out []*internaldto.OutputStrategyAveragePriceDto, err error)
	OpenClose(in *internaldto.InputStrategyOpenCloseDto) (out []*internaldto.OutputStrategyOpenCloseDto, err error)
}

func NewUsecaseStrategy(
	operationRepo mysql.OperationRepositoryInterface,
	operationHistoryRepo mysql.OperationHistoryRepositoryInterface,
	walletRepo mysql.WalletRepositoryInterface,
	parityRepo mysql.ParityRepositoryInterface,
	averageRepo mysql.AveragePriceRepositoryInterface,
) StrategyInteface {
	bybit := bybitSDK.NewBybitService(os.Getenv("BYBIT_API_KEY"), os.Getenv("BYBIT_SECRET_KEY"))

	return &strategy{
		ParityRepo:           parityRepo,
		OperationRepo:        operationRepo,
		OperationHistoryRepo: operationHistoryRepo,
		WalletRepo:           walletRepo,
		AverageRepo:          averageRepo,

		Bybit: bybit,
	}
}
