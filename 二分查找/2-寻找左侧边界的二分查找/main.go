package main

import "fmt"

// 寻找左侧边界的二分查找
func leftBound(nums []int, target int) int {
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
			right = mid - 1
		}
	}
	// 检查 left 越界情况
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func main() {
	nums := []int{1, 2, 2, 2, 4}
	res := leftBound(nums, 2)
	fmt.Println(res)
}
