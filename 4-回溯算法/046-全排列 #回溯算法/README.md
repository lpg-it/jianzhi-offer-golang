## 题目描述

给定一个**没有重复**数字的序列，返回其所有可能的全排列。

### 示例1

```go
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

## Golang实现

### 方法一



```go
package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	// 记录路径
	var track []int
	visited := make([]bool, len(nums))
	backtrack(nums, track, &res, visited)
	return res
}

func backtrack(nums []int, track []int, res *[][]int, visited []bool) {
	// 结束条件
	if len(nums) == len(track) {
		tmp:= make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 做选择
		if !visited[i] {
			visited[i] = true
			track = append(track, nums[i])
			// 进入下一层决策树
			backtrack(nums, track, res, visited)
			// 撤销选择
			track = track[:len(track)-1]
			visited[i] = false
		}
	}
}

func main() {
	nums := []int{5, 4, 6, 2}
	res := permute(nums)
	fmt.Println(res)
}
```

执行用时 : 0 ms, 在所有 Go 提交中击败了 100% 的用户

内存消耗 : 2.7 MB, 在所有 Go 提交中击败了 56.03% 的用户

最后附上本人的LeetCode题解链接：[https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/shu-zu-zhong-zhong-fu-de-shu-zi-by-lpgit/](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/shu-zu-zhong-zhong-fu-de-shu-zi-by-lpgit/)

## 李培冠博客

[lpgit.com](https://lpgit.com)
