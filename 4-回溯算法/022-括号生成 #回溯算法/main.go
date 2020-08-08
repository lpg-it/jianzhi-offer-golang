package main

import "fmt"

var res []string

func generateParenthesis(n int) []string {
	res = []string{}
	track := make([]string, 0, 2*n)
	backtrack(n, track, n, n)
	return res
}

func backtrack(n int, track []string, left int, right int) {
	// 如果左括号剩下的多, 说明左括号比右括号少, 不合法
	if left > right {
		return
	}
	// 数量小于0, 肯定不合法
	if left < 0 || right < 0 {
		return
	}

	// 所有括号用完, 得到一个合法的括号组合
	if left == 0 && right == 0 {
		str := ""
		for _, value := range track {
			str += value
		}
		res = append(res, str)
		return
	}
	// 做选择: 放一个左括号
	track = append(track, "(")
	// 回溯
	backtrack(n, track, left-1, right)
	// 撤销选择
	track = track[:len(track)-1]

	// 做选择: 放一个右括号
	track = append(track, ")")
	// 回溯
	backtrack(n, track, left, right-1)
	// 撤销选择
	track = track[:len(track)-1]
}

func main() {
	n := 3
	res := generateParenthesis(n)
	fmt.Println(res)

}
