package repositorydto

type InputOperationHistoryDto struct {
	OperationHistory        uint64
	Operation               uint64
	TransactionType         uint64
	CoinPrice               string
	CoinQuantity            string
	StablePrice             string
	StableQuantity          string
	Fee                     string
	OperationExchangeId     string
	OperationExchangeStatus uint64
}

type OutputOperationHistoryDto struct {
	OperationHistory        uint64
	Operation               uint64
	TransactionType         uint64
	CoinPrice               string
	CoinQuantity            string
	StablePrice             string
	StableQuantity          string
	Fee                     string
	OperationExchangeId     string
	OperationExchangeStatus uint64
}
