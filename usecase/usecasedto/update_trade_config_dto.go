package usecasedto

type InputUpdateTradeConfigDto struct {
	TradeConfig             uint64  `json:"trade_config"`
	User                    uint64  `json:"user"`
	Modality                uint64  `json:"modality"`
	Strategy                uint64  `json:"strategy"`
	StrategyVariant         uint64  `json:"strategy_variant"`
	Parity                  uint64  `json:"parity"`
	Exchange                uint64  `json:"exchange"`
	OperationQuantity       uint64  `json:"operation_quantity"`
	OperationAmount         float64 `json:"operation_amount"`
	DefaultProfitPercentage float64 `json:"default_profit_percentage"`
	WalletValueLimit        float64 `json:"wallet_value_limit"`
	Enabled                 bool    `json:"enabled"`
}

type OutputUpdateTradeConfigDto struct {
	TradeConfig             uint64  `json:"trade_config"`
	User                    uint64  `json:"user"`
	Modality                uint64  `json:"modality"`
	Strategy                uint64  `json:"strategy"`
	StrategyVariant         uint64  `json:"strategy_variant"`
	Parity                  uint64  `json:"parity"`
	Exchange                uint64  `json:"exchange"`
	OperationQuantity       uint64  `json:"operation_quantity"`
	OperationAmount         float64 `json:"operation_amount"`
	DefaultProfitPercentage float64 `json:"default_profit_percentage"`
	WalletValueLimit        float64 `json:"wallet_value_limit"`
	Enabled                 bool    `json:"enabled"`
}
