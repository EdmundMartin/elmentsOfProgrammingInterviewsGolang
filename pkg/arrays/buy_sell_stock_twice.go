package arrays

import (
	"math"
)

func sliceOfSize(defaultValue float64, size int) []float64 {
	retSlice := []float64{}
	for i := 0; i < size; i++ {
		retSlice = append(retSlice, defaultValue)
	}
	return retSlice
}

func BuySellStockTwice(prices []float64) float64 {
	maxTotalProfit, minPriceSoFar := 0.0, math.MaxFloat64
	firstBuySellProfits := sliceOfSize(0.0, len(prices))

	for idx, price := range prices {
		minPriceSoFar = math.Min(minPriceSoFar, price)
		maxTotalProfit = math.Max(maxTotalProfit, price-minPriceSoFar)
		firstBuySellProfits[idx] = maxTotalProfit
	}

	maxPriceSoFar := float64(math.MinInt32)
	for i := len(prices) - 1; i >= 0; i-- {
		maxPriceSoFar = math.Max(maxPriceSoFar, prices[i])
		runningProfit := maxPriceSoFar - prices[i] + firstBuySellProfits[i]
		maxTotalProfit = math.Max(maxTotalProfit, runningProfit)
	}
	return maxTotalProfit
}
