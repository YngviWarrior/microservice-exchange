package usecasedto

type Candle struct {
	Parity   uint64
	Exchange uint64
	Mts      uint64
	Open     float64
	Close    float64
	High     float64
	Low      float64
	Volume   float64
	Roc      float64
}

type InputListCandleLimitDto struct {
	Parity   uint64
	Exchange uint64
	Limit    uint64
}

type OutputListCandleLimitDto struct {
	Candles []*Candle
}
