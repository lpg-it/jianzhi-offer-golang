package main

import "fmt"

func quickSort(nums []int) []int {
	return _quickSort(nums, 0, len(nums)-1)
}

func _quickSort(nums []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(nums, left, right)
		_quickSort(nums, left, partitionIndex-1)
		_quickSort(nums, partitionIndex+1, right)
	}
	return nums
}

func partition(nums []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if nums[pivot] > nums[i] {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
	return index - 1
}

func main() {
	nums := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
	res := quickSort(nums)
	fmt.Println(res)
}
