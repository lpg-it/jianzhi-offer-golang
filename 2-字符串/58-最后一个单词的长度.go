package main

import "fmt"

func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res == 0 {
				continue
			} else {
				break
			}
		} else {
			res += 1
		}
	}
	return res
}

func main() {
	res := lengthOfLastWord("hello world ")
	fmt.Println(res)
}
