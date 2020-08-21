package main

import "fmt"

var res [][]int

func combine(n int, k int) [][]int {
	// 初始化 res
	res = [][]int{}
	// 记录路径
	track := make([]int, 0, k)
	backtrack(n, k, 1, track)
	return res
}

func backtrack(n int, k int, start int, track []int) {
	if k == len(track) {
		tmp := make([]int, k)
		copy(tmp, track)
		res = append(res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		// 做选择
		track = append(track, i)
		// 回溯
		backtrack(n, k, i+1, track)
		// 撤销选择
		track = track[:len(track)-1]
	}
}

func main() {
	n, k := 4, 2
	res := combine(n, k)
	fmt.Println(res)

}
