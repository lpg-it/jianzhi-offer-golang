package main

import "fmt"

// 计算两个数的最大值
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 方法一
/*func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	} else if numsLen == 1 {
		return nums[0]
	}
	// 定义一个切片dp, dp[i] 表示偷盗至 nums[i] 的最高金额
	dp := make([]int, numsLen)
	// 初始化为最小值
	for i := range dp {
		dp[i] = -1
	}

	// 初始化状态
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < numsLen; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	res := dp[0]
	for _, value := range dp {
		res = max(res, value)
	}
	return res
}*/

// 方法二
func rob(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	} else if numsLen == 1 {
		return nums[0]
	}
	nums[1] = max(nums[0], nums[1])
	for i := 2; i < numsLen; i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	res := nums[0]
	for _, value := range nums {
		res = max(res, value)
	}
	return res
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}
