package main

import "fmt"

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y%1000000007, (x+y)%1000000007
	}
	return x
}

func main() {
	n := 40
	fmt.Println(fib(n))
}
