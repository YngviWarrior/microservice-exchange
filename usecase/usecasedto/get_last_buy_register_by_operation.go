package usecasedto

type InputGetLastBuyRegisterByOperationDto struct {
	Operation uint64
}

type OutputGetLastBuyRegisterByOperationDto struct {
	CoinQuantity float64
	Fee          float64
}
