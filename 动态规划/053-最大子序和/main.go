package main

import (
	"fmt"
	"math"
)

/*func maxSubArray(nums []int) int {
	// 定义一个切片dp, dp[i] 表示以 nums[i] 结尾的连续子数组和最大的值
	dp := make([]int, len(nums))
	// 默认为最小数
	for i := range dp {
		dp[i] = -9999999
	}
	dp[0] = nums[0]
	for i, numsLen := 1, len(nums); i < numsLen; i++ {
		dp[i] = int(math.Max(float64(nums[i]), float64(dp[i-1]+nums[i])))
	}
	res := dp[0]
	for _, value := range dp {
		if value > res {
			res = value
		}
	}
	return res
}*/

func maxSubArray(nums []int) int {
	x, y := nums[0], 0
	res := nums[0]
	for i, numsLen := 1, len(nums); i < numsLen; i++ {
		y = int(math.Max(float64(nums[i]), float64(nums[i] + x)))
		x = y
		res = int(math.Max(float64(y), float64(res)))
	}
	return res
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
