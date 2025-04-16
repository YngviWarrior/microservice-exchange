package usecasedto

type InputUpdateAveragePriceDto struct {
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

type OutputUpdateAveragePriceDto struct{}
