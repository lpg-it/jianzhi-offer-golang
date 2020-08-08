package main

import (
	"fmt"
	"sort"
)

var res [][]int

func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	track := make([]int, 0)
	sort.Ints(candidates)
	backtrack(candidates, target, track, 0)
	return res
}

func backtrack(candidates []int, target int, track []int, start int) {
	// 如果 target 最终小于 0, 不合法
	if target < 0 {
		return
	}
	// 如果 target 最终等于 0, 结束条件
	if target == 0 {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}
	for i := start; i < len(candidates); i++ {
		// 去除重复项
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		// 做选择
		track = append(track, candidates[i])
		// 回溯
		backtrack(candidates, target-candidates[i], track, i+1)
		// 撤销选择
		track = track[:len(track)-1]
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	res := combinationSum2(candidates, target)
	fmt.Println(res)

}
