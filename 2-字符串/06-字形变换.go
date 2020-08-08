package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	res := make([]string, numRows)
	var i int
	for i = 0; i < numRows; i++ {
		res[i] += string(s[i])
	}
	i--
	flag := -1
	for j := numRows; j < len(s); j++ {
		i += flag
		res[i] += string(s[j])
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
	}
	var str string
	str = strings.Join(res, "")
	return str
}

func main() {
	s := "LEETCODEISHIRING"
	res := convert(s, 4)
	fmt.Println(res)
}
