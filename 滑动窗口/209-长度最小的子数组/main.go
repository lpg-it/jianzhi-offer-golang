package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(s int, nums []int) int {
	// 记录初始窗口
	left, right := 0, 0
	
	res := math.MaxInt32
	sum := 0
	for right < len(nums) {
		c := nums[right]
		right++
		sum += c

		// 判断是否应该左移
		for sum >= s {
			res = min(res, right-left)
			d := nums[left]
			left++
			sum -= d
		}
		
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min (x, y int) int {
	if x > y {
		return y
	}else {
		return x
	}
}


func main() {
	s := 7
	nums := []int{2,3,1,2,4,3}
	res := minSubArrayLen(s, nums)
	fmt.Println(res)

}
