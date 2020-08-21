package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	// 判断特殊情况
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	// dp[i][k][0] 表示第 i 天的第 1 笔交易没有股票时的最大利润;
	// dp[i][k-1][1] 表示第 i 天的第 2 笔交易持有股票时的最大利润

	// 初始化
	dpI10 := 0
	dpI20 := 0
	dpI11 := -prices[0]
	dpI21 := -prices[0]
	dpI0 := 0
	dpI1 := -prices[0]

	for i := 1; i < len(prices); i++ {
		dpI10 = max(dpI10, dpI11+prices[i])
		dpI11 = max(dpI11, -prices[i])
		// 今天没有股票:
		// 1. 昨天没有, 今天没有
		// 2. 昨天有, 今天卖了
		dpI20 = max(dpI20, dpI21+prices[i])
		// 今天有股票
		// 1. 昨天有, 今天有
		// 2. 昨天没有, 今天买了
		dpI21 = max(dpI21, dpI10-prices[i])

		// 一次交易
		dpI0 = max(dpI0, dpI1+prices[i])
		dpI1 = max(dpI1, -prices[i])
	}
	return max(dpI20, dpI0)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	prices := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	res := maxProfit(prices)
	fmt.Println(res)
}
