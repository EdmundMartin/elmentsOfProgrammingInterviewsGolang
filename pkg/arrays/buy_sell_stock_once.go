package arrays

import (
	"math"
)

func BuySellStockOnce(prices []float64) float64 {
	minPrice, maxProfit := float64(math.MaxInt32), 0.0
	for _, price := range prices {
		maxProfitToday := price - minPrice
		maxProfit = math.Max(maxProfit, maxProfitToday)
		minPrice = math.Min(minPrice, price)
	}
	return maxProfit
}
