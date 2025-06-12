package usecasedto

type InputUpdateOperationHistoryDto struct {
	OperationHistory        uint64 `json:"operation_history"`
	Operation               uint64 `json:"operation"`
	TransactionType         uint64 `json:"transaction_type"`
	CoinPrice               string `json:"coin_price"`
	CoinQuantity            string `json:"coin_quantity"`
	StablePrice             string `json:"stable_price"`
	StableQuantity          string `json:"stable_quantity"`
	Fee                     string `json:"fee"`
	OperationExchangeId     string `json:"operation_exchange_id"`
	OperationExchangeStatus uint64 `json:"operation_exchange_status"`
}

type OutputUpdateOperationHistoryDto struct {
	OperationHistory        uint64 `json:"operation_history"`
	Operation               uint64 `json:"operation"`
	TransactionType         uint64 `json:"transaction_type"`
	CoinPrice               string `json:"coin_price"`
	CoinQuantity            string `json:"coin_quantity"`
	StablePrice             string `json:"stable_price"`
	StableQuantity          string `json:"stable_quantity"`
	Fee                     string `json:"fee"`
	OperationExchangeId     string `json:"operation_exchange_id"`
	OperationExchangeStatus uint64 `json:"operation_exchange_status"`
}
