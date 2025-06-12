package usecasedto

type Candles struct {
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

type InputCreateCandlesDto struct {
	Candles []Candles
}

type OutputCreateCandlesDto struct{}
