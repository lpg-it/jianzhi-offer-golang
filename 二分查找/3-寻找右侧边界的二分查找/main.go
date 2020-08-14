package main

import "fmt"

func rightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}
	// 检查 right 越界情况
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

func main() {
	nums := []int{1, 2, 2, 4}
	res := rightBound(nums, 2)
	fmt.Println(res)
}
