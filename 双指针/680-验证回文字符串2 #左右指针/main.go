package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return isPalindrome(s[left:right]) || isPalindrome(s[left+1:right+1])
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "abcbfa"
	res := validPalindrome(s)
	fmt.Println(res)
}
