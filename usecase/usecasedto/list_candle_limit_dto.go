package usecasedto

type Candle struct {
	Parity   uint64
	Exchange uint64
	Mts      uint64
	Open     string
	Close    string
	High     string
	Low      string
	Volume   string
	Roc      string
}

type InputListCandleLimitDto struct {
	Parity   uint64
	Exchange uint64
	Limit    uint64
}

type OutputListCandleLimitDto struct {
	Candles []*Candle
}
