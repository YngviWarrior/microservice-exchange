package usecasedto

type InputGetLastBuyRegisterByOperationDto struct {
	Operation uint64
}

type OutputGetLastBuyRegisterByOperationDto struct {
	CoinQuantity string
	Fee          string
	Status       uint64
}
