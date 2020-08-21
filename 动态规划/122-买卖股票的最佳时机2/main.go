package main

func maxProfit(prices []int) int {
	// 判断特殊情况
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	// 因为不限制买卖次数, 所以不需要考虑这个因素
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	// 初始化 dp
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i, pricesLen := 1, len(prices); i < pricesLen; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}

// 优化版本
func maxProfit2(prices []int) int {
	// 判断特殊情况
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	// 因为不限制买卖次数, 所以不需要考虑这个因素

	// 初始化 dp
	dpI0 := 0
	dpI1 := -prices[0]

	for i, pricesLen := 1, len(prices); i < pricesLen; i++ {
		dpI0 = max(dpI0, dpI1+prices[i])
		dpI1 = max(dpI1, dpI0-prices[i])
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

}
