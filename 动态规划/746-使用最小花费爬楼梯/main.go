package main

import "fmt"

// 返回两个数的最小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 方法一
/*func minCostClimbingStairs(cost []int) int {
	costLen := len(cost)
	// 定义一个切片 dp, dp[i] 表示爬至第 i 个阶梯时, 花费的最少体力
	dp := make([]int, costLen)
	for i := range dp {
		dp[i] = 1<<31 - 1
	}
	// 初始化状态
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < costLen; i++ {
		dp[i] = min(dp[i-1]+cost[i], dp[i-2]+cost[i])
	}
	return min(dp[costLen-1], dp[costLen-2])
}*/

// 方法二
func minCostClimbingStairs(cost []int) int {
	costLen := len(cost)
	for i := 2; i < costLen; i++ {
		cost[i] = min(cost[i-1]+cost[i], cost[i-2]+cost[i])
	}
	return min(cost[costLen-1], cost[costLen-2])
}

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1} // [0, 1, 1, 1]
	fmt.Println(minCostClimbingStairs(cost))
}
