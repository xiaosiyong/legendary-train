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

func TestMaxVolumeInBag(t *testing.T) {
	fmt.Println(algorithm.MaxV)
	algorithm.MaxValueInBag(0, 0)
	fmt.Println(algorithm.MaxV)

}

func TestKnapsack(t *testing.T) {
	fmt.Println(algorithm.Knapsack(algorithm.Items, algorithm.Count, algorithm.Vol))
}
