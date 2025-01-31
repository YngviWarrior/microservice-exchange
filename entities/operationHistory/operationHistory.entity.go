package operationhistory

type OperationHistory struct {
	OperationHistory        int64   `json:"operation_history"`
	Operation               int64   `json:"operation"`
	TransactionType         int64   `json:"transaction_type"`
	CoinPrice               float64 `json:"coin_price"`
	CoinQuantity            float64 `json:"coin_quantity"`
	StablePrice             float64 `json:"stable_price"`
	StableQuantity          float64 `json:"stable_quantity"`
	OperationExchangeId     string  `json:"operation_exchange_id"`
	OperationExchangeStatus int64   `json:"operation_exchange_status"`
}

type OjoinOH struct {
	Operation      uint64  `json:"operation"`
	User           uint64  `json:"user_id"`
	Parity         uint64  `json:"parity"`
	ParitySymbol   string  `json:"parity_symbol"`
	Exchange       uint64  `json:"exchange"`
	MtsStart       int64   `json:"mts_start"`
	MtsFinish      int64   `json:"mts_finish"`
	OpenPrice      float64 `json:"open_price"`
	ClosePrice     float64 `json:"close_price"`
	InvestedAmount float64 `json:"invested_amount"`
	ProfitAmount   float64 `json:"profit_amount"`
	Profit         float64 `json:"profit"`
	Closed         bool    `json:"closed"`
	Enabled        bool    `json:"enabled"`
	Audit          bool    `json:"audit"`
	TimesCanceled  int64   `json:"times_canceled"`

	OperationHistory        int64   `json:"operation_history"`
	TransactionType         int64   `json:"transaction_type"`
	CoinPrice               float64 `json:"coin_price"`
	CoinQuantity            float64 `json:"coin_quantity"`
	StablePrice             float64 `json:"stable_price"`
	StableQuantity          float64 `json:"stable_quantity"`
	Fee                     float64 `json:"fee"`
	OperationExchangeId     string  `json:"operation_exchange_id"`
	OperationExchangeStatus int64   `json:"operation_exchange_status"`
}
