package main

import (
	"fmt"
	"sync"
)

// 两个 go 程轮流打印一个切片
func main() {
	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)
	ch1 <- true
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var i int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for ;  i < len(nums) && <-ch1; i++ {
			fmt.Println(nums[i])
			ch2 <- true
		}
		wg.Done()
	}()
	go func() {
		for ; <-ch2 && i < len(nums); i++ {
			fmt.Println(nums[i])
			ch1 <- true
		}
		wg.Done()
	}()
	wg.Wait()
}