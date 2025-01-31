package coin

type Coin struct {
	Coin   int64  `json:"coin"`
	Symbol string `json:"symbol"`
	Active bool   `json:"active"`
}
