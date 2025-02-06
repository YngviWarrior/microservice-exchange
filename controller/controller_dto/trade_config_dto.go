package controllerdto

type InputTradeConfigDto struct {
	Modality                *int64   `json:"modality" validate:"required"`
	Strategy                *int64   `json:"strategy" validate:"required"`
	StrategyVariant         *int64   `json:"strategy_variant" validate:"required"`
	Parity                  *int64   `json:"parity" validate:"required"`
	Exchange                *int64   `json:"exchange" validate:"required"`
	OperationQuantity       *int64   `json:"operation_quantity" validate:"required"`
	OperationAmount         *float64 `json:"operation_amount" validate:"required"`
	DefaultProfitPercentage *float64 `json:"default_profit_percentage" validate:"required"`
	WalletValueLimit        *float64 `json:"wallet_value_limit" validate:"required"`
	Enabled                 *bool    `json:"enabled" validate:"required"`
}
