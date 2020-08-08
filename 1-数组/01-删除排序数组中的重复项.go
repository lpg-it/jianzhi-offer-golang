package main

import "fmt"

/*
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
 */

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	preNum := nums[0]
	for _, v := range nums {
		if v != preNum {
			preNum = v
			nums[count] = v
			count += 1
		}
	}
	return count
}

func main() {
	var nums []int
	nums = append(nums, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4)
	res := removeDuplicates(nums)
	fmt.Println(res, nums)
}
