package averageprice

type AveragePrice struct {
	Parity               int64   `json:"parity"`
	Exchange             int64   `json:"exchange"`
	Day                  float64 `json:"day"`
	DayUpdateTimestamp   int64   `json:"day_update_timestamp"`
	Week                 float64 `json:"week"`
	WeekUpdateTimestamp  int64   `json:"week_update_timestamp"`
	Month                float64 `json:"month"`
	MonthUpdateTimestamp int64   `json:"month_update_timestamp"`
}
