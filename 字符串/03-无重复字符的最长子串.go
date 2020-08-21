package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	li := make([]byte, 1)
	li[0] = s[0]
	maxNum := 1
	for i:=1; i < len(s); i++{
		for j:=0; j < len(li); j++ {
			if s[i] == li[j] {
				// 已存在相同元素
				if maxNum < len(li) {
					maxNum = len(li)
				}
				li = li[j+1:]
				break
			}
		}

		li = append(li, s[i])
		if maxNum < len(li) {
			maxNum = len(li)
		}
	}
	return maxNum
}

func main() {
	s := "abcabcbb"
	num := lengthOfLongestSubstring(s)
	fmt.Println(num)
}
