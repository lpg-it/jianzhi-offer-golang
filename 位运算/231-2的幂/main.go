package main

import "fmt"

// 一个数如果是 2 的指数，那么它的二进制表示一定只含有一个 1：
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n-1)) == 0
}

func main() {
	n := 2
	res := isPowerOfTwo(n)
	fmt.Println(res)
}
