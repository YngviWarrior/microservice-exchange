package usecasedto

type InputTradeConfigDto struct {
	User                    uint64 `json:"user"`
	Modality                uint64 `json:"modality"`
	Strategy                uint64 `json:"strategy"`
	StrategyVariant         uint64 `json:"strategy_variant"`
	Parity                  uint64 `json:"parity"`
	Exchange                uint64 `json:"exchange"`
	OperationQuantity       uint64 `json:"operation_quantity"`
	OperationAmount         string `json:"operation_amount"`
	DefaultProfitPercentage string `json:"default_profit_percentage"`
	WalletValueLimit        string `json:"wallet_value_limit"`
	Enabled                 bool   `json:"enabled"`
}

type OutputTradeConfigDto struct {
	TradeConfig             uint64 `json:"trade_config"`
	User                    uint64 `json:"user"`
	Modality                uint64 `json:"modality"`
	Strategy                uint64 `json:"strategy"`
	StrategyEnabled         bool   `json:"strategy_enabled"`
	StrategyVariant         uint64 `json:"strategy_variant"`
	Parity                  uint64 `json:"parity"`
	Exchange                uint64 `json:"exchange"`
	OperationQuantity       uint64 `json:"operation_quantity"`
	OperationAmount         string `json:"operation_amount"`
	Enabled                 bool   `json:"enabled"`
	DefaultProfitPercentage string `json:"default_profit_percentage"`
	WalletValueLimit        string `json:"wallet_value_limit"`
	ModalityName            string `json:"modality_name"`
	StrategyName            string `json:"strategy_name"`
	StrategyVariantName     string `json:"strategy_variant_name"`
	StrategyVariantEnabled  bool   `json:"strategy_variant_enbled"`
	ParitySymbol            string `json:"symbol_name"`
	ExchangeName            string `json:"exchange_name"`
}
