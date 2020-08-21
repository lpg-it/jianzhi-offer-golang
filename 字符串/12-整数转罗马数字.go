package main

import (
	"fmt"
	"math"
	"strconv"
)

func intToRoman(num int) string {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	s := ""
	x := len(intSlice) - 1
	for ; num != 0; {
		i := len(strconv.Itoa(num)) - 1
		n := (num / int(math.Pow10(i))) * int(math.Pow10(i))
		for ; x >= 0; x-- {
			if n >= intSlice[x] {
				s += roman[x]
				num -= intSlice[x]
				break
			}
		}
	}
	return s
}

func main() {
	s := intToRoman(58)
	fmt.Println(s)
}
