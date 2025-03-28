package usecasedto

type InputTradeConfigDto struct {
	User                    int64   `json:"user"`
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
	User                    int64   `json:"user"`
	Modality                int64   `json:"modality"`
	Strategy                int64   `json:"strategy"`
	StrategyEnabled         bool    `json:"strategy_enabled"`
	StrategyVariant         int64   `json:"strategy_variant"`
	Parity                  int64   `json:"parity"`
	Exchange                int64   `json:"exchange"`
	OperationQuantity       int64   `json:"operation_quantity"`
	OperationAmount         float64 `json:"operation_amount"`
	Enabled                 bool    `json:"enabled"`
	DefaultProfitPercentage float64 `json:"default_profit_percentage"`
	WalletValueLimit        float64 `json:"wallet_value_limit"`
	ModalityName            string  `json:"modality_name"`
	StrategyName            string  `json:"strategy_name"`
	StrategyVariantName     string  `json:"strategy_variant_name"`
	StrategyVariantEnabled  bool    `json:"strategy_variant_enbled"`
	ParitySymbol            string  `json:"symbol_name"`
	ExchangeName            string  `json:"exchange_name"`
}
