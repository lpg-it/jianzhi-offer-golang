## 题目描述

在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

### 示例1

现有矩阵 matrix 如下：

```go
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
```
给定 target = 5，返回 true。

给定 target = 20，返回 false。

## Golang实现

### 方法一

暴力解法，没啥可说的，外层控制行，内层控制列，循环判断即可。

```go
package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != target {
				continue
			} else {
				return true
			}
		}
	}
	return false
}

func main() {
	n := 1
	var matrix [][]int
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}
	matrix[0] = []int{1, 4, 7, 11, 15}
	b := findNumberIn2DArray(matrix, 20)
	fmt.Println(b)
}
```

执行用时 : 32 ms, 在所有 Go 提交中击败了 76.59% 的用户

内存消耗 : 6.2 MB, 在所有 Go 提交中击败了 100% 的用户

### 方法二

参考了别人的思路。

由于这个二维数组的每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。所以：

- **左下角数据**：所在**列**最大数据，所在**行**最小数据

- **右上角数据**：所在**行**最大数据，所在**列**最小数据

故，我们可以先对左下角数据进行对比，如果该数据大于 target，则直接去掉该数据所在行；如果该数据小于 target，则直接去掉该数据所在列；如果相等，直接返回。

```go
package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for row, col := len(matrix)-1, 0; row >= 0 && col < len(matrix[0]);  {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}
	return false
}

func main() {
	n := 1
	var matrix [][]int
	for i := 0; i < n; i++ {
		matrix = append(matrix, make([]int, n))
	}
	matrix[0] = []int{1, 4, 7, 11, 15}
	b := findNumberIn2DArray(matrix, 20)
	fmt.Println(b)
}
```

执行用时 : 32 ms, 在所有 Go 提交中击败了 76.59% 的用户

内存消耗 : 6.3 MB, 在所有 Go 提交中击败了 100% 的用户

最后附上本人的LeetCode题解链接：[https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/solution/er-wei-shu-zu-zhong-de-cha-zhao-liang-chong-jie-fa/](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/solution/er-wei-shu-zu-zhong-de-cha-zhao-liang-chong-jie-fa/)

## 李培冠博客

[lpgit.com](https://lpgit.com)
