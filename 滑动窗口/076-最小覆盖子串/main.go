package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	// 记录窗口中含有 需要凑齐 的字符
	window := map[byte]int{}
	// 需要凑齐的字符
	need := map[byte]int{}
	for i := range t {
		need[t[i]]++
	}

	// 初始化窗口两端
	left, right := 0, 0

	// 表示窗口中满足 need 条件的字符个数, 如果 valid 和 len(need) 相同, 则说明窗口满足条件
	valid := 0

	// 记录最小覆盖子串的起始索引及长度
	start, length := 0, math.MaxInt32

	for right < len(s) {
		// 开始滑动
		// c 是被移入窗口的字符
		c := s[right]
		// 右移窗口
		right++
		/* 进行窗口内数据的一系列更新 */
		if _, ok := need[c]; ok {
			// need 中存在该数据
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right - left < length {
				start = left
				length = right - left
			}
			// d 是要移出窗口的字符
			d := s[left]
			// 左移窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				window[d]--
				if window[d] < need[d] {
					valid--
				}
			}
		}
	}
	// 返回最小覆盖子串
	if length == math.MaxInt32 {
		return ""
	}
	return s[start:start+length]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	res := minWindow(s, t)
	fmt.Println(res)
}
