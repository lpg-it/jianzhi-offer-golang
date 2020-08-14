package main

func characterReplacement(s string, k int) int {
	// 初始化窗口两端
	left, right := 0, 0

	//
	window := map[byte]int{}

	// 存储重复子串的最长结果
	res := 0

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := window[c]; !ok {
			window[c]++
			k--
		} else {
			window[c]++
		}


		for k == 0 {

		}

	}

}

func main() {
	
}
