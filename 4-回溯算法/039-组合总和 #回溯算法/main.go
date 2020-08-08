package main

import "fmt"

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	track := make([]int, 0)
	backtrack(candidates, target, track, 0)
	return res
}

func backtrack(candidates []int, target int, track []int, start int) {
	// target 小于 0, 错误
	if 0 > target {
		return
	}
	// 判断结束条件
	if 0 == target {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		// 做选择
		track = append(track, candidates[i])
		// 回溯
		backtrack(candidates, target-candidates[i], track, i)
		// 撤销选择
		track = track[:len(track)-1]
	}
}

func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	res := combinationSum(candidates, target)
	fmt.Println(res)
}
