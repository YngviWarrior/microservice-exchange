package repositorydto

type InputOperationHistoryDto struct{}
type OutputOperationHistoryDto struct {
	OperationHistory        uint64
	Operation               uint64
	TransactionType         uint64
	CoinPrice               float64
	CoinQuantity            float64
	StablePrice             float64
	StableQuantity          float64
	Fee                     float64
	OperationExchangeId     string
	OperationExchangeStatus uint64
}
