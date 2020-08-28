package main

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	return num&(num-1) == 0 && (num&1431655765) == num
}

func main() {

}
