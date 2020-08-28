package main

import "fmt"

func mergeSort(nums []int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return nums
	}

	middle := numsLen / 2
	left := nums[:middle]
	right := nums[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) != 0 {
		result = append(result, left...)
	}

	if len(right) != 0 {
		result = append(result, right...)
	}

	return result
}

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		nums[i] = tmp
	}
	res := mergeSort(nums)
	fmt.Println(res)
}
