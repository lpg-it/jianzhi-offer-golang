package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {

	for i, j := 0, 1; j < len(s);  {




		if s[i] == s[j] {
			i++

		}

	}
}

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
}
