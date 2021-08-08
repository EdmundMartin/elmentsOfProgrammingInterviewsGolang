package arrays

import "testing"

func TestBuySellStockOnce(t *testing.T) {
	prices := []float64{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}
	maxProfit := BuySellStockOnce(prices)
	if maxProfit != 30.0 {
		t.Errorf("expected 30.0, got: %f", maxProfit)
	}
}
