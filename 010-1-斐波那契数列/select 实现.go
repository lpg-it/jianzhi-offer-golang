package main

import (
	"fmt"
	"runtime"
)

func fibSelect(ch <- chan int, quit <- chan bool) int {
	for  {
		select {
		case num := <- ch:
			fmt.Println(num)
		case <- quit:
			runtime.Goexit()
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go fibSelect(ch, quit)

	x, y := 0, 1
	for i := 0; i < 30; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}
