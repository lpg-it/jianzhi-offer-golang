package main

import "fmt"

func hammingDistance(x int, y int) int {
	z := x ^ y
	// 转换成了求位 1 的个数问题
	count := 0
	for z != 0 {
		z  = z & (z - 1)
		count++
	}
	return count
}

func main() {
	x, y := 1, 4
	res := hammingDistance(x, y)
	fmt.Println(res)
}
