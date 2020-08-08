package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	i, j := 0, len(height) - 1
	var res float64
	for ;i < j;{
		currArea := float64(j - i) * math.Min(float64(height[i]), float64(height[j]))
		res = math.Max(res, currArea)
		if height[i] > height[j]{
			j--
		} else {
			i++
		}
	}
	return int(res)
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(height)
	fmt.Println(res)
}
