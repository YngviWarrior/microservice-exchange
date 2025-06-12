package usecasedto

type InputGetLastPriceDto struct {
	Parity   uint64
	Exchange uint64
}

type OutputGetLastPriceDto struct {
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
