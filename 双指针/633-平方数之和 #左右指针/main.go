package main

import "math"

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		if left*left+right*right > c {
			right--
		} else if left*left+right*right < c {
			left++
		} else if left*left+right*right == c {
			return true
		}
	}
	return false
}

func main() {

}
