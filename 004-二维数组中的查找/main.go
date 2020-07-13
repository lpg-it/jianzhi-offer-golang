package main

import (
	"fmt"
)

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

/*// 暴力解法
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
}*/

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
