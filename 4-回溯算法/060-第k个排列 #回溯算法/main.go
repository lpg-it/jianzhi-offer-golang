package main

import (
	"fmt"
	"strconv"
)

var res []string

func getPermutation(n int, k int) string {
	res = []string{}
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	track := make([]string, 0)
	visited := make([]bool, len(nums))
	backtrack(nums, track, visited, k,0)
	return res[k-1]
}

func backtrack(nums []int, track []string, visited []bool, k int,start int) {
	if len(res) == k {
		return
	}
	if len(track) == len(nums) {
		str := ""
		for _, s := range track {
			str += s
		}
		res = append(res, str)
		return
	}
	for i := start; i < len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			// 做选择
			track = append(track, strconv.Itoa(nums[i]))
			// 回溯
			backtrack(nums, track, visited, k, start)
			// 撤销选择
			track = track[:len(track)-1]
			visited[i] = false
		}
	}
}

func main() {
	n, k := 8, 29805
	res := getPermutation(n, k)
	fmt.Println(res)
}
