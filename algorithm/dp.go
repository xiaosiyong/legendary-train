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

func FbN1(n int) int {
	if n < 3 {
		return 1
	}
	return FbN1(n-1) + FbN1(n-2)
}

func FbN2(n int) int {
	pre := 1
	if n < 2 {
		return pre
	}
	cur := 2
	for i := 2; i < n; i++ {
		temp := cur
		cur += pre
		pre = temp
	}
	return cur
}

func Knapsack(weight []int, n, w int) int {
	states := make([][]bool, n)
	for i := range states {
		states[i] = make([]bool, w+1)
	}
	//先处理第一个
	states[0][0] = true
	if weight[0] < w {
		states[0][weight[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ { //当前不放
			if states[i-1][j] {
				states[i][j] = true
			}
		}
		for j := 0; j <= w-weight[i]; j++ { //当前放进去
			if states[i-1][j] {
				states[i][weight[i]+j] = true
			}
		}
	}
	for i := w; i >= 0; i-- {
		if states[n-1][i] {
			return i
		}
	}
	return 0
}

func Knapsack2(weight []int, n, w int) int {
	states := make([]bool, w+1)
	//先处理第一个
	states[0] = true
	for i := 1; i < n; i++ {
		for j := w - weight[i]; j >= 0; j-- {
			if states[j] {
				states[j+weight[i]] = true
			}
		}
	}
	for i := w; i >= 0; i-- {
		if states[i] {
			return i
		}
	}
	return 0
}

func KnapspackWithValue(weight, value []int, n, w int) int {
	states := make([][]int, n)
	for i := range states {
		states[i] = make([]int, w+1)
		for j := 0; j <= w; j++ {
			states[i][j] = -1
		}
	}
	//先处理第一个
	states[0][0] = 0
	if weight[0] < w {
		states[0][weight[0]] = value[0]
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ { //当前不放
			if states[i-1][j] >= 0 {
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= w-weight[i]; j++ { //当前放进去
			if states[i-1][j] >= 0 {
				v := states[i-1][j] + value[i]
				if v > states[i][weight[i]+j] {
					states[i][weight[i]+j] = v
				}
			}
		}
	}
	max := -1
	for i := w; i >= 0; i-- {
		if states[n-1][i] > max {
			max = states[n-1][i]
		}
	}
	return max
}
