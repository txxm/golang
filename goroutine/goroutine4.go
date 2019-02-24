package main

import (
	"fmt"
	"runtime"
)

/* 不加runtime.Gosched(),goroutine未执行完可能主goroutine就退出，可能只输出World */
func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		fmt.Println("Hello")
	}()

	runtime.Gosched()
	fmt.Println("World")

	function()
}

/* 让出时间片后，在10000次循环中，主goroutine可能继续执行，导致提前退出 */
func function() {
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println("i =", i)
		}
	}()

	runtime.Gosched()
	fmt.Println("golang")
}
