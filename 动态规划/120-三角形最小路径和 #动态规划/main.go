package main

import (
	"fmt"
	"math"
)

// 方法一
/*func minimumTotal(triangle [][]int) int {
	triangleRow := len(triangle)
	// 定义一个二维数组dp, dp[i][j] 表示包含 dp[i][j] 在内的最小路径和
	dp := make([][]int, triangleRow)
	for i := range dp {
		dp[i] = make([]int, i+1)
		for j := range dp[i] {
			dp[i][j] = 1<<31 - 1
		}
	}
	// 定义初始状态
	dp[0][0] = triangle[0][0]
	for i := 1; i < triangleRow; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				// 第 0 列, 只能从自己头顶的元素下来
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				// 每一行中的最后一列, 只能从自己的左上角元素下来
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				// 中间元素, 可以从自己的头顶或者右上元素下来
				dp[i][j] = int(math.Min(float64(dp[i-1][j-1]), float64(dp[i-1][j]))) + triangle[i][j]
			}
		}
	}
	res := dp[triangleRow-1][0]
	for i := range dp[triangleRow-1] {
		if res > dp[triangleRow-1][i] {
			res = dp[triangleRow-1][i]
		}
	}
	return res
}*/

// 方法二
func minimumTotal(triangle [][]int) int {
	triangleRow := len(triangle)
	for i := 1; i < triangleRow; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				// 第 0 列, 只能从自己头顶的元素下来
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				// 每一行中的最后一列, 只能从自己的左上角元素下来
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			} else {
				// 中间元素, 可以从自己的头顶或者右上元素下来
				triangle[i][j] = int(math.Min(float64(triangle[i-1][j-1]), float64(triangle[i-1][j]))) + triangle[i][j]
			}
		}
	}
	res := triangle[triangleRow-1][0]
	for i := range triangle[triangleRow-1] {
		if res > triangle[triangleRow-1][i] {
			res = triangle[triangleRow-1][i]
		}
	}
	return res
}


func main() {
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}
