package arrays

import "testing"

func TestBuySellStockTwice(t *testing.T) {
	prices := []float64{12, 11, 13, 9, 12, 8, 14, 13, 15}
	maxProfit := BuySellStockTwice(prices)
	if maxProfit != 10.0 {
		t.Errorf("expected 10.0, got %f", maxProfit)
	}
}
