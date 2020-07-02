package main

import (
	"fmt"
	"sort"
)

/*// 方法一
func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, value := range nums {
		if _, ok := m[value]; ok {
			// 存在该数值
			return value
		} else {
			// 不存在
			m[value] = 1
		}
	}
	return -1
}*/

// 方法二
func findRepeatNumber(nums []int) int {
	// 对整型切片进行排序
	sort.Ints(nums)

	// 遍历切片, 与下一个数字重复就返回
	for i, numsSize := 0, len(nums); i < numsSize-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	num := findRepeatNumber(nums)
	fmt.Println("重复的一个数字为: ", num)
}
