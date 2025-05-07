package usecasedto

type InputGetAvgPriceByParityExchangeDto struct {
	Parity   uint64
	Exchange uint64
}

type OutputGetAvgPriceByParityExchangeDto struct {
	Parity               uint64
	Exchange             uint64
	Day                  float64
	DayRoc               float64
	DayUpdateTimestamp   uint64
	Week                 float64
	WeekRoc              float64
	WeekUpdateTimestamp  uint64
	Month                float64
	MonthRoc             float64
	MonthUpdateTimestamp uint64
}
