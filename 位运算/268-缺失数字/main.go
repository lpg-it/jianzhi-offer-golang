package main

func missingNumber(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	res := len(nums) * (len(nums) + 1) / 2
	return res - sum
}

func main() {

}
