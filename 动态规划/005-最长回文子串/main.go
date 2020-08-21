package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	sLen := len(s)
	if sLen == 1 {
		return true
	}
	for i := 0; i < sLen/2; i++ {
		if s[i] != s[sLen-1-i] {
			return false
		}
	}
	return true
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func longestPalindrome(s string) string {
	// 定义一个切片 dp, dp[i] 表示至 s[i] 的最长回文字串的个数
	dp := make([]string, len(s))
	dp[0] = string(s[0])
	for i := 1; i < len(s); i++ {
		if isPalindrome(s[:i]) {
			// 是回文
			dp[i] = max(dp[i-1]+1, dp[i])
		} else {
			// 不是回文
			dp[i] = dp[i-1]
		}




	}

	fmt.Println(isPalindrome(s))
	return " "
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
