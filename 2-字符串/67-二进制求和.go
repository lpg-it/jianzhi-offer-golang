package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	aSize := len(a)
	bSize := len(b)
	if aSize < bSize {
		for i := 0; i < bSize - aSize; i++ {
			a = "0" + a
		}
	} else if aSize > bSize {
		for i := 0; i < aSize - bSize; i++ {
			b = "0" + b
		}
	}

	fmt.Println(a, b)

	flag := -1
	binSum := ""

	for i := len(a) - 1; i >= 0; i-- {
		ai, _ := strconv.Atoi(string(a[i]))
		bi, _ := strconv.Atoi(string(b[i]))
		var c int
		if flag == 1 {
			c = ai + bi + 1
		} else {
			c = ai + bi
		}

		if c == 0 {
			flag = -1
			binSum = strconv.Itoa(c) + binSum
		}else if c == 1 {
			flag = -1
			binSum = strconv.Itoa(c) + binSum
		}else if c == 2 {
			flag = 1
			binSum = "0" + binSum
		}else if c == 3 {
			flag = 1
			binSum = "1" + binSum
		}
		if i == 0 && flag == 1 {
			binSum = "1" + binSum
		}
	}
	return binSum
}

func main() {
	a := "100"
	b := "110010"
	s := addBinary(a, b)
	fmt.Println(s)
}
