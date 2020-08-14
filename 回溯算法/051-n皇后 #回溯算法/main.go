package main

import "fmt"

var res [][]string

func solveNQueens(n int) [][]string {
	// 初始化结果
	res = [][]string{}
	// 记录路径
	track := make([][]string, 0, n)
	// 初始化 track
	for i := 0; i < n; i++ {
		tmp := make([]string, 0, n)
		for j := 0; j < n; j++ {
			tmp = append(tmp, ".")
		}
		track = append(track, tmp)
	}
	backtrack(track, 0)
	return res
}

func backtrack(track [][]string, row int) {
	// 触发结束条件
	if len(track) == row {
		tmp := make([]string, 0, len(track))
		// 将每行的选择结果改为字符串
		for _, value := range track {
			str := ""
			for _, s := range value {
				str += s
			}
			tmp = append(tmp, str)
		}
		res = append(res, tmp)
		return
	}

	n := len(track[row])
	for col := 0; col < n; col++ {
		// 判断此处是否可以选择皇后(同行, 同列, 对角不能存在多个皇后)
		if isValid(track, row, col) {
			// 选择皇后
			track[row][col] = "Q"
			// 进行下一行决策选择
			backtrack(track, row+1)
			// 撤销选择
			track[row][col] = "."
		}
	}
}
func isValid(track [][]string, row int, col int) bool {
	n := len(track)
	// 判断同一列
	for i := 0; i < n; i++ {
		if track[i][col] == "Q" {
			return false
		}
	}

	// 左上角
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if track[r][c] == "Q" {
			return false
		}
	}

	// 右上角
	for r, c := row-1, col+1; r >= 0 && c < n; r, c = r-1, c+1 {
		if track[r][c] == "Q" {
			return false
		}
	}
	return true
}
func main() {
	n := 1
	res := solveNQueens(n)
	fmt.Println(res)
}
