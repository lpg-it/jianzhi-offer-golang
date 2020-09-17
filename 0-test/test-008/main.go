package main

import "fmt"

// 快速排序
func quickSort(nums []int) []int {
	return _quickSort(nums, 0, len(nums)-1)
}

func _quickSort(nums []int, left, right int) []int {
	if left < right {
		partitionIndex, rightIndex := partition(nums, left, right)
		_quickSort(nums, left, partitionIndex-1)
		_quickSort(nums, rightIndex, right)
	}
	return nums
}

func partition(nums []int, left, right int) (int, int) {
	pivot := left
	leftIndex := pivot + 1
	for i := leftIndex; i <= right; {
		if nums[pivot] > nums[i] {
			nums[i], nums[leftIndex] = nums[leftIndex], nums[i]
			leftIndex++
			i++
		} else if nums[pivot] < nums[i] {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			i++
		}
	}
	fmt.Println(right)
	nums[pivot], nums[leftIndex-1] = nums[leftIndex-1], nums[pivot]
	return leftIndex - 1, right
}

func main() {

	nums := []int{7, 2, 3, 1, 7, 9, 8, 10, 4, 12}
	res := quickSort(nums)
	fmt.Println(res)
}
