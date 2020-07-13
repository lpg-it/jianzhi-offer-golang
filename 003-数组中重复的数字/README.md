## 题目描述

找出数组中重复的数字。

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

### 示例1

```go
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3 
```

## Golang实现

### 方法一

由于只需要找出数组中任意一个重复的数字，因此遍历数组，遇到重复的数字即返回。为了判断一个数字是否重复遇到，使用 map 存储已经遇到的数字，如果遇到的一个数字已经在 map 中，则当前的数字是重复数字。


```go
package main

import (
	"fmt"
)

func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, value := range nums {
		if _, ok := m[value]; ok {
			// 存在该数值
			return value
		} else {
			// 不存在
			m[value] = 1
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	num := findRepeatNumber(nums)
	fmt.Println("重复的一个数字为: ", num)
}
```

执行用时 : 76 ms, 在所有 Go 提交中击败了 7.58% 的用户

内存消耗 : 9.1 MB, 在所有 Go 提交中击败了 100% 的用户

### 方法二

排序，看前后元素是否相等

```go
package main

import (
	"fmt"
	"sort"
)

func findRepeatNumber(nums []int) int {
	// 对整型切片进行排序
	sort.Ints(nums)

	// 遍历切片, 与下一个数字重复就返回
	for i, numsSize := 0, len(nums); i < numsSize-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	num := findRepeatNumber(nums)
	fmt.Println("重复的一个数字为: ", num)
}

```

执行用时 : 56 ms, 在所有 Go 提交中击败了 22.35% 的用户

内存消耗 : 6.9 MB, 在所有 Go 提交中击败了 100% 的用户

最后附上本人的LeetCode题解链接：[https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/shu-zu-zhong-zhong-fu-de-shu-zi-by-lpgit/](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/solution/shu-zu-zhong-zhong-fu-de-shu-zi-by-lpgit/)

## 李培冠博客

[lpgit.com](https://lpgit.com)
