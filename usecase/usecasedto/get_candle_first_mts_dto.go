package usecasedto

type InputGetCandleFirstMtsDto struct {
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

type OutputGetCandleFirstMtsDto struct {
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
