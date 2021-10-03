package helper

import (
	"math"

	"github.com/google/uuid"
)

const USDCode = "USD"

func GetUUID() string {
	return uuid.New().String()
}

func GetConvertedToUSDAmount(rate, amount int64) int64 {
	return int64(math.Round((100 / float64(rate)) * float64(amount)))
}

func GetConvertedFromUSDAmount(rate, amount int64) int64 {
	return int64(math.Round((float64(rate) / 100) * float64(amount)))
}

func GetConvertedCurrency(rateFrom, rateTo, amount int64) int64 {
	return GetConvertedFromUSDAmount(rateTo, GetConvertedToUSDAmount(rateFrom, amount))
}
