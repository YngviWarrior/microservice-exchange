package usecasedto

type InputGetOperationHistoryDto struct {
	OrderExchangeId uint64
}

type OutputGetOperationHistoryDto struct {
	OperationHistory        uint64  `json:"operation_history"`
	Operation               uint64  `json:"operation"`
	TransactionType         uint64  `json:"transaction_type"`
	CoinPrice               float64 `json:"coin_price"`
	CoinQuantity            float64 `json:"coin_quantity"`
	StablePrice             float64 `json:"stable_price"`
	StableQuantity          float64 `json:"stable_quantity"`
	Fee                     float64 `json:"fee"`
	OperationExchangeId     string  `json:"operation_exchange_id"`
	OperationExchangeStatus uint64  `json:"operation_exchange_status"`
}
