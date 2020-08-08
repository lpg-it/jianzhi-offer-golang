package main

import "fmt"

/* 198 - 打家劫舍 */
func rob(nums []int) int {
	numsLen := len(nums)
	// dp[i] 表示到 i 号房屋为止, 最高的金额
	dp := make([]int, numsLen)
	// 初始化
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < numsLen; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[numsLen-1]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {
	var num int
	fmt.Scan(&num)
	var tmp int
	nums := make([]int, num, num)
	for i :=0; i<num; i++ {
		fmt.Scan(&tmp)
		nums[i] = tmp
	}
	res := rob(nums)
	fmt.Println(res)
}
