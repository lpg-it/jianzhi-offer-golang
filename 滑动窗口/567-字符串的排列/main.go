package main

func checkInclusion(s1 string, s2 string) bool {
	// 记录窗口中含有 需要凑齐 的字符
	window := map[byte]int{}
	// 需要凑齐的字符
	needs := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		needs[s1[i]]++
	}
	// 初始化窗口两端
	left, right := 0, 0
	// 表示窗口中满足 needs 的个数, 如果 valid 与 len(needs) 相同, 则表示窗口满足条件
	valid := 0

	for right < len(s2) {
		// 开始滑动, c 是要移入的字符
		c := s2[right]
		// 右移窗口
		right++
		// 更新数据
		if _, ok := needs[c]; ok {
			window[c]++
			if window[c] == needs[c] {
				valid++
			}
		}

		// 判断左侧窗口是否收缩
		for right-left >= len(s1) {
			// 判断是否找到了合法的子串
			if valid == len(needs) {
				return true
			}
			// 将要移出的字符
			d := s2[left]
			left++
			// 进行窗口的更新
			if _, ok := needs[d]; ok {
				if window[d] == needs[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

func main() {

}
