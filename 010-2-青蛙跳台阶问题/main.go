package main

import "fmt"

func numWays(n int) int {
	if n == 0 || n == 1{return 1}
	x, y := 1, 2
	for i := 0; i < n-1; i++ {
		x, y = y%1000000007, (x + y)%1000000007
	}
	return x
}

func main() {
	fmt.Println(numWays(2))
}
