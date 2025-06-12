package usecasedto

type InputUpdateOperationDto struct {
	Operation       uint64 `json:"operation"`
	User            uint64 `json:"user_id"`
	Parity          uint64 `json:"parity"`
	Exchange        uint64 `json:"exchange"`
	Strategy        uint64 `json:"strategy"`
	StrategyVariant uint64 `json:"strategy_variant"`
	MtsStart        uint64 `json:"mts_start"`
	MtsFinish       uint64 `json:"mts_finish"`
	OpenPrice       string `json:"open_price"`
	ClosePrice      string `json:"close_price"`
	InvestedAmount  string `json:"invested_amount"`
	ProfitAmount    string `json:"profit_amount"`
	Profit          string `json:"profit"`
	Closed          bool   `json:"closed"`
	Audit           bool   `json:"audit"`
	Enabled         bool   `json:"enabled"`
	InTransaction   bool   `json:"in_transaction"`
}

type OutputUpdateOperationDto struct {
	Operation       uint64 `json:"operation"`
	User            uint64 `json:"user_id"`
	Parity          uint64 `json:"parity"`
	Exchange        uint64 `json:"exchange"`
	Strategy        uint64 `json:"strategy"`
	StrategyVariant uint64 `json:"strategy_variant"`
	MtsStart        uint64 `json:"mts_start"`
	MtsFinish       uint64 `json:"mts_finish"`
	OpenPrice       string `json:"open_price"`
	ClosePrice      string `json:"close_price"`
	InvestedAmount  string `json:"invested_amount"`
	ProfitAmount    string `json:"profit_amount"`
	Profit          string `json:"profit"`
	Closed          bool   `json:"closed"`
	Audit           bool   `json:"audit"`
	Enabled         bool   `json:"enabled"`
	InTransaction   bool   `json:"in_transaction"`
}
