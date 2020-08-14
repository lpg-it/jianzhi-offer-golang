package main

import "fmt"

func findAnagrams(s string, p string) []int {
	// 需要凑齐的字符
	needs := map[byte]int{}
	for i := range p {
		needs[p[i]]++
	}

	// 记录窗口中 需要凑齐 的字符
	window := map[byte]int{}

	// 初始化窗口两端
	left, right := 0, 0

	// valid 表示窗口中满足 needs 中字符的个数
	valid := 0

	// 记录满足条件的起始索引
	var res []int

	for right < len(s) {
		// 开始滑动, c 是将要移入的字符
		c := s[right]
		right++
		// 更新数据
		if _, ok := needs[c]; ok {
			window[c]++
			if window[c] == needs[c] {
				valid++
			}
		}

		// 判断是否需要左移
		for right-left == len(p) {
			// 添加起始索引
			if valid == len(needs) {
				res = append(res, left)
			}

			// 将要移出的字符
			d := s[left]
			left++
			// 更新窗口中的数据
			if _, ok := needs[d]; ok {
				if window[d] == needs[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	res := findAnagrams(s, p)
	fmt.Println(res)
}
