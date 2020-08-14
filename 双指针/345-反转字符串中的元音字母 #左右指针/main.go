package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	vowels := map[byte]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0, 'A': 0, 'E': 0, 'I': 0, 'O': 0, 'U': 0}
	left, right := 0, len(s)-1
	sli := strings.Split(s, "")
	for left < right {
		if _, ok := vowels[s[left]]; !ok {
			// 不是元音字母
			left++
			continue
		}
		if _, ok := vowels[s[right]]; !ok {
			// 不是元音字母
			right--
			continue
		}
		sli[left], sli[right] = sli[right], sli[left]
		left++
		right--
	}
	s = strings.Join(sli, "")
	return s
}

func main() {
	s := "hello"
	res := reverseVowels(s)
	fmt.Println(res)
}
