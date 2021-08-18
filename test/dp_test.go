package test

import (
	"fmt"
	"legendaryTrain/algorithm"
	"testing"
)

func TestCoinChanges(t *testing.T) {
	result1 := algorithm.CoinChanges([]int{1, 2, 3, 4, 5}, 15)
	fmt.Println(result1)

	maxPrice := algorithm.StockProfit([]int{11, 8, 9, 16, 20})
	fmt.Println(maxPrice)
}
