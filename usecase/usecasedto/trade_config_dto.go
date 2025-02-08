package usecasedto

type InputTradeConfigDto struct {
	Modality                int64   `json:"modality"`
	Strategy                int64   `json:"strategy"`
	StrategyVariant         int64   `json:"strategy_variant"`
	Parity                  int64   `json:"parity"`
	Exchange                int64   `json:"exchange"`
	OperationQuantity       int64   `json:"operation_quantity"`
	OperationAmount         float64 `json:"operation_amount"`
	DefaultProfitPercentage float64 `json:"default_profit_percentage"`
	WalletValueLimit        float64 `json:"wallet_value_limit"`
	Enabled                 bool    `json:"enabled"`
}

type OutputTradeConfigDto struct {
	TradeConfig             int64   `json:"trade_config"`
	Modality                int64   `json:"modality"`
	Strategy                int64   `json:"strategy"`
	StrategyVariant         int64   `json:"strategy_variant"`
	Parity                  int64   `json:"parity"`
	Exchange                int64   `json:"exchange"`
	OperationQuantity       int64   `json:"operation_quantity"`
	OperationAmount         float64 `json:"operation_amount"`
	DefaultProfitPercentage float64 `json:"default_profit_percentage"`
	WalletValueLimit        float64 `json:"wallet_value_limit"`
	Enabled                 bool    `json:"enabled"`
}
