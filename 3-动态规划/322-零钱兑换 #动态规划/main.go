package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	coinCount := make([]int, amount + 1, amount + 1)
	for i := range coinCount {
		coinCount[i] = amount + 1
	}
	coinCount[0] = 0
	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i < coin{continue}
			coinCount[i] = int(math.Min(float64(coinCount[i]), 1 + float64(coinCount[i-coin])))
		}
	}
	if coinCount[amount] == amount+1{
		return -1
	}
	return coinCount[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
