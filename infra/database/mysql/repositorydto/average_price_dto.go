package repositorydto

type InputAveragePriceDto struct {
	Parity               uint64
	Exchange             uint64
	Day                  string
	DayUpdateTimestamp   uint64
	Week                 string
	WeekUpdateTimestamp  uint64
	Month                string
	MonthUpdateTimestamp uint64
	Sma                  string
	SmaUpdateTimestamp   uint64
	DayRoc               string
	WeekRoc              string
	MonthRoc             string
	SmaRoc               string
}

type OutputAveragePriceDto struct {
	Parity               uint64
	Exchange             uint64
	Day                  string
	DayUpdateTimestamp   uint64
	Week                 string
	WeekUpdateTimestamp  uint64
	Month                string
	MonthUpdateTimestamp uint64
	Sma                  string
	SmaUpdateTimestamp   uint64
	DayRoc               string
	WeekRoc              string
	MonthRoc             string
	SmaRoc               string
}
