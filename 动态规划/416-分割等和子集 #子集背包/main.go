package main

import (
	"fmt"
	"math"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if math.Mod(float64(sum), float64(2)) != 0 {
		// 和为奇数, 肯定不成立
		return false
	}
	sum = sum / 2
	// dp[i][j] = true 表示截止到 i 为止, 存在和为 j 的子集
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	// 初始化 dp, dp[][0] 应该都为 true
	for i := range dp {
		dp[i][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum; j++ {
			if nums[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][sum]
}

func main() {
	nums := []int{1, 5, 11, 5}
	res := canPartition(nums)
	fmt.Println(res)
}
