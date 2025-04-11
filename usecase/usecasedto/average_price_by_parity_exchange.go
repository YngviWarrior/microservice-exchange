package usecasedto

type InputGetAvgPriceByParityExchangeDto struct {
	Parity   uint64
	Exchange uint64
}

type OutputGetAvgPriceByParityExchangeDto struct {
	Parity uint64
	Symbol string
	Active bool
}
