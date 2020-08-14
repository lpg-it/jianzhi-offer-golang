package main

import "fmt"

func searchRange(nums []int, target int) []int {
	res := make([]int, 2)
	for i := range res {
		res[i] = -1
	}
	if len(nums) == 0 {
		return res
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}
	res[0] = left
	if left >= len(nums) || nums[left] != target {
		res[0] = -1
	}
	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}
	res[1] = right
	if right < 0 || nums[right] != target {
		res[1] = -1
	}
	return res
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	res := searchRange(nums, 8)
	fmt.Println(res)

}
