package repositorydto

type InputTradeConfigDto struct {
	TradeConfig             uint64
	User                    uint64
	Modality                uint64
	Strategy                uint64
	StrategyVariant         uint64
	Parity                  uint64
	Exchange                uint64
	OperationQuantity       uint64
	OperationAmount         float64
	DefaultProfitPercentage float64
	WalletValueLimit        float64
	Enabled                 bool
}

type OutputTradeConfigDto struct {
	TradeConfig             uint64
	User                    uint64
	Modality                uint64
	Strategy                uint64
	StrategyEnabled         bool
	StrategyVariant         uint64
	Parity                  uint64
	Exchange                uint64
	OperationQuantity       uint64
	OperationAmount         float64
	Enabled                 bool
	DefaultProfitPercentage float64
	WalletValueLimit        float64
	ModalityName            string
	StrategyName            string
	StrategyVariantName     string
	StrategyVariantEnabled  bool
	ParitySymbol            string
	ExchangeName            string
}
