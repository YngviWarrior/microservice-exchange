package usecasedto

type InputCreateOperationDto struct {
	Operation       uint64  `json:"operation"`
	User            uint64  `json:"user_id"`
	Parity          uint64  `json:"parity"`
	Exchange        uint64  `json:"exchange"`
	Strategy        uint64  `json:"strategy"`
	StrategyVariant uint64  `json:"strategy_variant"`
	MtsStart        uint64  `json:"mts_start"`
	MtsFinish       uint64  `json:"mts_finish"`
	OpenPrice       float64 `json:"open_price"`
	ClosePrice      float64 `json:"close_price"`
	InvestedAmount  float64 `json:"invested_amount"`
	ProfitAmount    float64 `json:"profit_amount"`
	Profit          float64 `json:"profit"`
	Closed          bool    `json:"closed"`
	Audit           bool    `json:"audit"`
	Enabled         bool    `json:"enabled"`
	InTransaction   bool    `json:"in_transaction"`
}

type OutputCreateOperationDto struct {
	Operation       uint64  `json:"operation"`
	User            uint64  `json:"user_id"`
	Parity          uint64  `json:"parity"`
	Exchange        uint64  `json:"exchange"`
	Strategy        uint64  `json:"strategy"`
	StrategyVariant uint64  `json:"strategy_variant"`
	MtsStart        uint64  `json:"mts_start"`
	MtsFinish       uint64  `json:"mts_finish"`
	OpenPrice       float64 `json:"open_price"`
	ClosePrice      float64 `json:"close_price"`
	InvestedAmount  float64 `json:"invested_amount"`
	ProfitAmount    float64 `json:"profit_amount"`
	Profit          float64 `json:"profit"`
	Closed          bool    `json:"closed"`
	Audit           bool    `json:"audit"`
	Enabled         bool    `json:"enabled"`
	InTransaction   bool    `json:"in_transaction"`
}
