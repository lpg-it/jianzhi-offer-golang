package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// 创建二维切片
	dp := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, 2)
	}
	// 特殊处理
	dp[0][0] = 0
	dp[0][1] = -float64(prices[0])

	for i := 1; i < n; i++ {
		dp[i][0] = math.Max(dp[i - 1][0], dp[i - 1][1] + float64(prices[i]))
		dp[i][1] = math.Max(dp[i - 1][1], -float64(prices[i]))
	}
	return int(dp[n - 1][0])
}

func main() {
	s := make([]int, 0)
	s = append(s, 7, 1, 5, 3, 6, 4)
	i := maxProfit(s)
	fmt.Println(i)
}

// 0ms
// 4.3MB
