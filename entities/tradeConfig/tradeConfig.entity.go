package tradeconfig

type TradeConfig struct {
	TradeConfig            uint64  `json:"trade_config"`
	Modality               uint64  `json:"modality"`
	Strategy               uint64  `json:"strategy"`
	StrategyEnabled        bool    `json:"strategy_enabled"`
	StrategyVariant        uint64  `json:"strategy_variant"`
	Parity                 uint64  `json:"parity"`
	Exchange               uint64  `json:"exchange"`
	OperationQuantity      int64   `json:"operation_quantity"`
	OperationAmount        float64 `json:"operation_amount"`
	Enabled                bool    `json:"enabled"`
	ProfitPercentage       float64 `json:"profit_percentage"`
	WalletValueLimit       float64 `json:"wallet_value_limit"`
	ModalityName           string  `json:"modality_name"`
	StrategyName           string  `json:"strategy_name"`
	StrategyVariantName    string  `json:"strategy_variant_name"`
	StrategyVariantEnabled bool    `json:"strategy_variant_enbled"`
	ParitySymbol           string  `json:"symbol_name"`
	ExchangeName           string  `json:"exchange_name"`
}
