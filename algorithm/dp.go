package algorithm

import "math"

func CoinChanges(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt8
		for _, coin := range coins {
			if i >= coin && dp[i-coin] != math.MaxInt8 {
				dp[i] = getMin(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[amount] == math.MaxInt8 {
		return -1
	}
	return dp[amount]
}

func getMin(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func StockProfit(price []int) int {
	maxProfit := math.MinInt8
	minPrice := price[0]
	for i := 1; i < len(price); i++ {
		if price[i] < minPrice {
			minPrice = price[i]
		}
		temp := price[i] - minPrice
		if temp > maxProfit {
			maxProfit = temp
		}
	}
	return maxProfit
}
