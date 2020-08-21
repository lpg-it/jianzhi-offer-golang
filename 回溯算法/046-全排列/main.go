package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	// 记录路径
	var track []int
	visited := make([]bool, len(nums))
	backtrack(nums, track, &res, visited)
	return res
}

func backtrack(nums []int, track []int, res *[][]int, visited []bool) {
	// 结束条件
	if len(nums) == len(track) {
		tmp:= make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 做选择
		if !visited[i] {
			visited[i] = true
			track = append(track, nums[i])
			// 进入下一层决策树
			backtrack(nums, track, res, visited)
			// 撤销选择
			track = track[:len(track)-1]
			visited[i] = false
		}
	}
}

func main() {
	nums := []int{5, 4, 6, 2}
	res := permute(nums)
	fmt.Println(res)
}
