package main

import (
	"fmt"
	"math"
)

func lengthOfLIS(nums []int) int {
	// 定义一个切片dp, dp[i] 表示以 nums[i] 结尾的最长上升子序列的长度
	dp := make([]int, len(nums))
	// 初始化全为1
	for i := range dp {
		dp[i] = 1
	}
	for i, numsLen := 0, len(nums); i < numsLen; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = int(math.Max(float64(dp[i]), 1+float64(dp[j])))
			}
		}
	}
	maxLen := 0
	for _, value := range dp {
		if value > maxLen {
			maxLen = value
		}
	}
	return maxLen
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
