package main

import (
	"fmt"
	"math"
)

// 动态规划
func maxSubArray(nums []int) int {
	// 定义一个切片dp, dp[i] 表示以 nums[i] 结尾的最大和
	dp := make([]int, len(nums), len(nums))
	// 初始化每个元素为最小值
	for i := range dp {
		dp[i] = -101
	}
	dp[0] = nums[0]
	// 表示最大的和
	res := -101
	for i, numsLen := 1, len(nums); i < numsLen; i++ {
		dp[i] = int(math.Max(float64(nums[i]), float64(dp[i-1]+nums[i])))
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
