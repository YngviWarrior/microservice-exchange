package usecasedto

type InputAvgPricesDto struct {
	To   uint64
	From uint64
}

type OutputAvgPricesDto struct {
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
