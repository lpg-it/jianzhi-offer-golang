## 题目描述

输入一个整型数组，数组里有正数也有负数。数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。

要求时间复杂度为O(n)。

### 示例1

```go
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```

## Golang实现

动态规划实现最简单，直接梭哈。把数组定义对了就好办了。

```go
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
```

执行用时 : 24 ms, 在所有 Go 提交中击败了 92.81% 的用户

内存消耗 : 6.2 MB, 在所有 Go 提交中击败了 100% 的用户

最后附上本人的LeetCode题解链接：[https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/solution/lian-xu-zi-shu-zu-de-zui-da-he-dong-tai-gui-hua-3/](https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/solution/lian-xu-zi-shu-zu-de-zui-da-he-dong-tai-gui-hua-3/)

## 李培冠博客

[lpgit.com](https://lpgit.com)
