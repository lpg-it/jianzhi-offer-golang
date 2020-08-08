package main

import "fmt"

var res [][]int

func subsets(nums []int) [][]int {
	// 初始化 res
	res = [][]int{[]int{}}
	track := make([]int, 0)
	backtrack(nums, track, 0)
	return res
}

func backtrack(nums []int, track []int, start int) {
	// 要从 start 开始递增
	for i := start; i < len(nums); i++ {
		// 做选择
		track = append(track, nums[i])
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		// 回溯
		backtrack(nums, track, i+1)
		// 撤销回溯
		track = track[:len(track)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
}
