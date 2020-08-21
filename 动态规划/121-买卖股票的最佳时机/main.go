package main

import "fmt"

func maxProfit(prices []int) int {
	// 判断特殊情况
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	// dp[i][0] 表示第 i 天没有股票时的最大利润;
	// dp[i][1] 表示第 i 天持有股票时的最大利润
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// 初始化 dp
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i, pricesLen := 1, len(prices); i < pricesLen; i++ {
		// 第 i 天没有股票, 有两种情况:
		// 前一天没有股票, 今天仍然没有股票
		// 前一天有股票, 今天卖出, 没有股票
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])

		// 第 i 天有股票, 有两种情况:
		// 前一天有股票, 今天仍然有股票
		// 前一天没有股票, 今天买入, 有股票
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[len(prices)-1][0]
}

// 优化版本
func maxProfit2(prices []int) int {
	// 判断特殊情况
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	// dpI0 表示第 i 天没有股票时的最大利润;
	// dpI1 表示第 i 天持有股票时的最大利润

	// 初始化 dp
	dpI0 := 0
	dpI1 := -prices[0]

	for i, pricesLen := 1, len(prices); i < pricesLen; i++ {
		// 第 i 天没有股票, 有两种情况:
		// 前一天没有股票, 今天仍然没有股票
		// 前一天有股票, 今天卖出, 没有股票
		dpI0 = max(dpI0, dpI1+prices[i])

		// 第 i 天有股票, 有两种情况:
		// 前一天有股票, 今天仍然有股票
		// 前一天没有股票, 今天买入, 有股票
		dpI1 = max(dpI1, -prices[i])
	}
	return dpI0
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)
	fmt.Println(res)
}
