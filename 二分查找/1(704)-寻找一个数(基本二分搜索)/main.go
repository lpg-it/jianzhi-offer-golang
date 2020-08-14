package main

import "fmt"

// 搜搜一个数, 如果存在, 返回其索引, 否则返回 -1
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target{
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}


func main(){
	var x int
	fmt.Scan(&x)
	tmp := 0
	nums := make([]int, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&tmp)
		nums[i] = tmp
	}
	var target int
	fmt.Scan(&target)
	res := binarySearch(nums, target)
	fmt.Println(res)
}
