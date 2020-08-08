package main

import (
	"fmt"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 方法一
/*func minPathSum(grid [][]int) int {
	gridRow, gridCol := len(grid), len(grid[0])
	// 定义一个二维切片 dp, dp[i][j] 表示经过 grid[i][j] 的最小路径和
	dp := make([][]int, gridRow)
	// 初始化成最大值
	for i := range dp {
		dp[i] = make([]int, gridCol)
		for j := 0; j < gridCol; j++ {
			dp[i][j] = 1<<31-1
		}
	}
	fmt.Println(dp)
	for i := 0; i < gridRow; i++ {
		for j := 0; j < gridCol; j++ {
			// 最初状态初始化
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
				continue
			}
			// 特殊处理
			if i == 0 && j != 0 {
				// 第 0 行, 右边的值只能从左边来
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0{
				// 第 0 列, 下面的值只能从上面来
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[gridRow-1][gridCol-1]
}*/

// 方法二
func minPathSum(grid [][]int) int {
	gridRow, gridCol := len(grid), len(grid[0])
	// 直接使用grid, grid[i][j] 被修改成最小路径和
	for i := 0; i < gridRow; i++ {
		for j := 0; j < gridCol; j++ {
			// 特殊处理
			if i == 0 && j != 0 {
				// 第 0 行, 右边的值只能从左边来
				grid[i][j] = grid[i][j-1] + grid[i][j]
			} else if j == 0 && i != 0{
				// 第 0 列, 下面的值只能从上面来
				grid[i][j] = grid[i-1][j] + grid[i][j]
			} else if i != 0 && j != 0 {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			}
		}
	}
	return grid[gridRow-1][gridCol-1]
}

func main() {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}

	fmt.Println(minPathSum(grid))
}
