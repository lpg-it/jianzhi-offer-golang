package main

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else if nums[left]+nums[right] == target {
			res[0] = left+1
			res[1] = right+1
			break
		}
	}
	return res
}

func main() {

}
