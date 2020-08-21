package main

import "fmt"

func singleNumber(nums []int) int {
	return (sum(set(nums))*3 - sum(nums)) / 2
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func set(nums []int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	for key := range m {
		res = append(res, key)
	}
	return res
}

func main() {
	nums := []int{2, 2, 3, 2}
	res := singleNumber(nums)
	fmt.Println(res)
}
