package main

import (
	"fmt"
	"sort"
)

var res [][]int

func subsetsWithDup(nums []int) [][]int {
	// 为了更好的判断重复元素, 先对 nums 进行排序
	sort.Ints(nums)
	// 初始化 res, 必有空集
	res = [][]int{[]int{}}

	// 记录选择路径
	var track []int
	// 回溯
	backtrack(nums, track, 0)
	return res
}

func backtrack(nums []int, track []int, start int) {
	for i := start; i < len(nums); i++ {
		// 判断是否为重复元素
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		track = append(track, nums[i])
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		// 回溯
		backtrack(nums, track, i+1)
		track = track[:len(track)-1]
	}
}

func main() {
	nums := []int{1, 2, 2}
	res := subsetsWithDup(nums)
	fmt.Println(res)
}
