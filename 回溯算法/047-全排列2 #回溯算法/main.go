package main

import (
	"fmt"
	"sort"
)

var res [][]int

func permuteUnique(nums []int) [][]int {
	res = [][]int{}
	track := make([]int, 0)
	sort.Ints(nums)
	visited := make([]bool, len(nums))
	backtrack(nums, track, visited, 0)
	return res
}

func backtrack(nums []int, track []int, visited []bool, start int) {
	// 结束条件
	if len(nums) == len(track) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] {
			continue
		}
		if !visited[i] {
			visited[i] = true
			// 做选择
			track = append(track, nums[i])
			// 回溯
			backtrack(nums, track, visited, i+1)
			// 撤销选择
			track = track[:len(track)-1]
			visited[i] = false
		}
	}
}

func main() {
	nums := []int{1, 1, 2}
	res := permuteUnique(nums)
	fmt.Println(res)
}
