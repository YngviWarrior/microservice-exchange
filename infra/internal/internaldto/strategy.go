package internaldto

type TradeConfig struct {
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
	UserName                string
	ModalityName            string
	StrategyName            string
	StrategyVariantName     string
	StrategyVariantEnabled  bool
	SymbolName              string
	ExchangeName            string
	ParitySymbol            string
}

type InputStrategyAveragePriceDto struct {
	ClosePrice  float64
	TradeConfig *TradeConfig
}

type OutputStrategyAveragePriceDto struct{}

type InputStrategyOpenCloseDto struct {
	ClosePrice  float64
	TradeConfig *TradeConfig
}

type OutputStrategyOpenCloseDto struct{}
