package main

import "fmt"

func climbStairs(n int) int {
	x, y := 1, 2
	for i := 1; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	n := 1
	fmt.Println(climbStairs(n))
}
