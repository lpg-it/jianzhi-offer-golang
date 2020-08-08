package main

import "fmt"

/* golang 二维切片初始化 */
func main() {
	// 初始化方法1
	row, column := 3, 4
	var answer [][]int
	for i := 0; i < row; i++ {
		inline := make([]int, column)
		answer = append(answer, inline)
	}
	fmt.Println(answer)

	// 初始化方法2
	answer1 := make([][]int, row)
	for i := range answer1 {
		answer1[i] = make([]int, column)
	}
	fmt.Println(answer1)
}