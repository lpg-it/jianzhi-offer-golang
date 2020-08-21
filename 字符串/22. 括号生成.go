package main

import "fmt"

func generateParenthesis(n int) []string {
	m := map[string]int{
		"(": n,
		")": n,
	}
	for i := 0; i < n * 2; i++ {

	}
}

func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}
