package repositorydto

type InputTradeConfigDto struct {
	User                    int64
	Modality                int64
	Strategy                int64
	StrategyVariant         int64
	Parity                  int64
	Exchange                int64
	OperationQuantity       int64
	OperationAmount         float64
	DefaultProfitPercentage float64
	WalletValueLimit        float64
	Enabled                 bool
}

type OutputTradeConfigDto struct {
	TradeConfig             int64
	User                    int64
	Modality                int64
	Strategy                int64
	StrategyEnabled         bool
	StrategyVariant         int64
	Parity                  int64
	Exchange                int64
	OperationQuantity       int64
	OperationAmount         float64
	Enabled                 bool
	DefaultProfitPercentage float64
	WalletValueLimit        float64
	UserName                string
	ModalityName            string
	StrategyName            string
	StrategyVariantName     string
	StrategyVariantEnabled  bool
	ParitySymbol            string
	ExchangeName            string
}
