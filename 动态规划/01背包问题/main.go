package main

import "fmt"

/*
给你一个可装载重量为 w 的背包和 n 个物品, 每个物品有重量和价值两个属性

其中第 i 个物品的重量为 wt[i], 价值为 val[i],

现在让你用这个背包装物品, 最多能装的价值是多少
*/
func knapsack(w, n int, wt, val []int) int {
	// dp[i][j] 表示 截止到 i 这个物品时, 重量为 j 的时候的最大价值
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, w+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-wt[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], val[i-1]+dp[i-1][j-wt[i-1]])
			}
		}
	}
	return dp[n][w]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	n, w := 3, 4
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}
	res := knapsack(w, n, wt, val)
	fmt.Println(res)
}
