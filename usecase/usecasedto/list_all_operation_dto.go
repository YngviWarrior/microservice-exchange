package usecasedto

type InputListOperationAllDto struct{}

type OutputListOperationAllDto struct {
	Operation           uint64 `json:"operation"`
	User                uint64 `json:"user_id"`
	Parity              uint64 `json:"parity"`
	Exchange            uint64 `json:"exchange"`
	Strategy            uint64 `json:"strategy"`
	StrategyName        string `json:"strategy_name"`
	StrategyVariant     uint64 `json:"strategy_variant"`
	StrategyVariantName string `json:"strategy_variant_name"`
	MtsStart            uint64 `json:"mts_start"`
	MtsFinish           uint64 `json:"mts_finish"`
	OpenPrice           string `json:"open_price"`
	ClosePrice          string `json:"close_price"`
	InvestedAmount      string `json:"invested_amount"`
	ProfitAmount        string `json:"profit_amount"`
	Profit              string `json:"profit"`
	Closed              bool   `json:"closed"`
	Audit               bool   `json:"audit"`
	Enabled             bool   `json:"enabled"`

	// OperationMetaDataFastTrade sql.NullInt64
	// MinimumPrice               sql.NullString
	// MaximumPrice               sql.NullString
	// LastPrice                  sql.NullString
	// IsReceding                 sql.NullBool
}
