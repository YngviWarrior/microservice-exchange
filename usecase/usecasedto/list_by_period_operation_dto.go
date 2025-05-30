package usecasedto

type InputListOperationByPeriodDto struct {
	MtsStart uint64
	MtsEnd   uint64
}

type OutputListOperationByPeriodDto struct {
	Operation           uint64  `json:"operation"`
	User                uint64  `json:"user_id"`
	Parity              uint64  `json:"parity"`
	Exchange            uint64  `json:"exchange"`
	Strategy            uint64  `json:"strategy"`
	StrategyName        string  `json:"strategy_name"`
	StrategyVariant     uint64  `json:"strategy_variant"`
	StrategyVariantName string  `json:"strategy_variant_name"`
	MtsStart            uint64  `json:"mts_start"`
	MtsFinish           uint64  `json:"mts_finish"`
	OpenPrice           float64 `json:"open_price"`
	ClosePrice          float64 `json:"close_price"`
	InvestedAmount      float64 `json:"invested_amount"`
	ProfitAmount        float64 `json:"profit_amount"`
	Profit              float64 `json:"profit"`
	Closed              bool    `json:"closed"`
	Audit               bool    `json:"audit"`
	Enabled             bool    `json:"enabled"`

	// OperationMetaDataFastTrade sql.NullInt64
	// MinimumPrice               sql.NullFloat64
	// MaximumPrice               sql.NullFloat64
	// LastPrice                  sql.NullFloat64
	// IsReceding                 sql.NullBool
}
