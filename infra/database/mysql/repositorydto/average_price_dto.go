package repositorydto

type InputAveragePriceDto struct {
	Parity               uint64
	Exchange             uint64
	Day                  float64
	DayUpdateTimestamp   uint64
	Week                 float64
	WeekUpdateTimestamp  uint64
	Month                float64
	MonthUpdateTimestamp uint64
	Sma                  float64
	SmaUpdateTimestamp   uint64
	DayRoc               float64
	WeekRoc              float64
	MonthRoc             float64
	SmaRoc               float64
}

type OutputAveragePriceDto struct {
	Parity               uint64
	Exchange             uint64
	Day                  float64
	DayUpdateTimestamp   uint64
	Week                 float64
	WeekUpdateTimestamp  uint64
	Month                float64
	MonthUpdateTimestamp uint64
	Sma                  float64
	SmaUpdateTimestamp   uint64
	DayRoc               float64
	WeekRoc              float64
	MonthRoc             float64
	SmaRoc               float64
}
