package main

func lengthOfLongestSubstring(s string) int {
	// 记录窗口中每一个字符的个数
	window := map[byte]int{}
	// 初始化窗口两端
	left, right := 0, 0
	var res int

	for right < len(s) {
		// 开始滑动, c 是将要移入的字符
		c := s[right]
		right++
		window[c]++

		// 判断是否应该左移
		for window[c] >= 2 {
			// 开始左移, d 是将要移出的字符
			d := s[left]
			left++
			window[d]--
		}
		res = max(res, right-left)

	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {

}
