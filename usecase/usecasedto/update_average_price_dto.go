package usecasedto

type InputUpdateAveragePriceDto struct {
	Parity               uint64
	Exchange             uint64
	Day                  string
	DayRoc               string
	DayUpdateTimestamp   uint64
	Week                 string
	WeekRoc              string
	WeekUpdateTimestamp  uint64
	Month                string
	MonthRoc             string
	MonthUpdateTimestamp uint64
}

type OutputUpdateAveragePriceDto struct{}
