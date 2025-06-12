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
	OperationAmount         string
	DefaultProfitPercentage string
	WalletValueLimit        string
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
	OperationAmount         string
	Enabled                 bool
	DefaultProfitPercentage string
	WalletValueLimit        string
	ModalityName            string
	StrategyName            string
	StrategyVariantName     string
	StrategyVariantEnabled  bool
	ParitySymbol            string
	ExchangeName            string
}
