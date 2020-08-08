package main

import "fmt"

/* 全排列, 回溯算法 */

// 输入一组不重复的数字, 返回他们的全排列
func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	// 记录路径
	var track []int
	// 记录是否被选择
	visited := make([]bool, len(nums))
	backtrack(nums, track, &res, visited)
	return res
}

// 路径: 记录在 track 中
// 选择列表: nums 中不存在于 track 的哪些元素
// 结束列表: nums 中的元素全都在 track 中出现
func backtrack(nums []int, track []int, res *[][]int, visited []bool) {
	// 触发结束条件
	if len(nums) == len(track) {
		*res = append(*res, track)
		return
	}
	for i, num := range nums {
		// 排除不合法的选择
		if !visited[i] {
			visited[i] = true
			// 做选择
			track = append(track, num)
			// 进入下一层决策树
			backtrack(nums, track, res, visited)
			// 取消选择
			track = track[:len(track)-1]
			visited[i] = false
		}
	}
}

func main() {
	var num int
	fmt.Scan(&num)
	nums := make([]int, 0, num)
	var a int
	for i := 0; i < num; i++ {
		fmt.Scan(&a)
		nums = append(nums, a)
	}
	fmt.Println(nums)
	res := permute(nums)
	fmt.Println(res)
}
