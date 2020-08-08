package main

import "fmt"

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := m[s[len(s) - 1]]
	for i := len(s) - 1; i > 0; i-- {
		preNum := m[s[i-1]]
		if m[s[i]] > preNum {
			n -= preNum
		}else {
			n += preNum
		}
	}
	return n
}

func main() {
	s := "CIV"
	n := romanToInt(s)
	fmt.Println(n)
}
