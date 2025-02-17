package repositorydto

type InputOperationDto struct {
	Operation       uint64
	User            uint64
	Parity          uint64
	Exchange        uint64
	Strategy        uint64
	StrategyVariant uint64
	MtsStart        uint64
	MtsFinish       uint64
	OpenPrice       float64
	ClosePrice      float64
	InvestedAmount  float64
	ProfitAmount    float64
	Profit          float64
	Closed          bool
	Audit           bool
	Enabled         bool
	TimesCanceled   float64
}

type OutputOperationDto struct {
	Operation       uint64
	User            uint64
	Parity          uint64
	Exchange        uint64
	Strategy        uint64
	StrategyVariant uint64
	MtsStart        uint64
	MtsFinish       uint64
	OpenPrice       float64
	ClosePrice      float64
	InvestedAmount  float64
	ProfitAmount    float64
	Profit          float64
	Closed          bool
	Audit           bool
	Enabled         bool
	TimesCanceled   float64
}
