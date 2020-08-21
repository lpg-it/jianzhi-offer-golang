package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	str2 := "11"
	res := ""
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}
	x := 1
	for i := 2; i < n; i++ {
		for j := 0; j < len(str2); j++ {
			y := string(str2[j])
			afterStr := string(str2[j + 1])
			if string(str2[j]) != afterStr || j == len(str2) - 1{
				res += strconv.Itoa(x)
				res += y
				x = 1
				if j == len(str2) - 1 {
					break
				}
			}else {
				x += 1
				y = string(str2[j])
			}
		}
		str2 = res
		res = ""
	}
	return str2
}

func main() {
	s := countAndSay(4)
	fmt.Println(s)

}
