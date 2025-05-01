package internaldto

type BuyCoinParams struct {
	Operation  uint64
	ClosePrice float64
	Exchange   uint64
	OpAmount   string
	Symbol     string
	OrderType  string
	OpFee      float64
}

type SellCoinParams struct {
	Operation    uint64
	ClosePrice   float64
	CoinQuantity float64
	Exchange     uint64
	OpAmount     float64
	OpFee        float64
	Symbol       string
	OrderType    string
}
