package main

import "fmt"

func foo(a [][]int) int {
	dp := make([][]int, len(a)+2)
	for i := 0; i < len(a)+2; i++ {
		dp[i] = make([]int, len(a[0])+2)
	}

	dp[1][1] = 1
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
		dp[i][len(dp[0])-1] = 1
	}
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = 1
		dp[len(dp)-1][i] = 1
	}
	fmt.Println(dp)


	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(a[0]); j++ {

			dp[i][j] = dp[i-1][j] + dp[i][j-1] + dp[i+1][j] + dp[i][j+1]
		}
	}
	fmt.Println(dp)
	return 1



}

func main() {
	b := [][]int{
		[]int{8, 4, 1},
		[]int{6, 5, 2},
	}
	res := foo(b)
	fmt.Println(res)





}
