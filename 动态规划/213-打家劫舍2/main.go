package main

import (
	"fmt"
)

func foo(nums []int) int {
	// dp[i] 表示到目前为止, 所能偷窃的最高金额
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	return max(foo(nums[1:]), foo(nums[:len(nums)-1]))
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	nums := []int{2, 3, 2}
	res := rob(nums)
	fmt.Println(res)
}
