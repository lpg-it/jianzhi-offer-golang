package main

// 两个相同的数异或为 0, 一个与 0 异或的数还是这个数本身
func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}

func main() {
	
}
