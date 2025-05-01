package utils

import (
	"fmt"
	"strconv"
)

func Truncate(some float64, decimal int64) float64 {
	var decimals string = "10"

	for i := 0; i < int(decimal); i++ {
		decimals = fmt.Sprintf("%v0", decimals)
	}

	decimalsToFloat, _ := strconv.ParseFloat(decimals, 64)

	return float64(int(some*decimalsToFloat)) / decimalsToFloat
}
