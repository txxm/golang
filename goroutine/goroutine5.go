package main

import (
	"fmt"
	"time"
	"runtime"
)

/* runtime.Gosched()让出当前goroutine时间片，让其他的goroutine执行 */
func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		fmt.Printf("Hello ")
	}()
	runtime.Gosched()
	fmt.Println("World")

	function()
}

func function() {
	/* runtime.Goexit()终止当前goroutine，不会影响其他的goroutine */
	go func() {
		/* 在终止当前goroutine之前，先执行未执行的defer */
		defer fmt.Println("Current goroutine exit.")
		for i := 0; i < 5; i++ {
			if i >= 2 {
				runtime.Goexit()
			} else {
				fmt.Println("i =", i)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("golang")
}
