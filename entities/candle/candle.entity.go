package candle

type Candle struct {
	Parity   int64
	Exchange int64
	Mts      int64
	Open     float64
	Close    float64
	High     float64
	Low      float64
	Volume   float64
	Roc      float64
}
