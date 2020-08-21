package main

import "fmt"

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num-1)
		count++
	}
	return count
}

func main() {
	var num uint32 = 00000000000000000000000000001011
	res := hammingWeight(num)
	fmt.Println(res)
}
