package main

import (
	"fmt"
	"sync"
)

/* 顺序输出 */
func main() {
	var go_sync sync.WaitGroup
	for i := 0; i < 5; i++ {
		go_sync.Add(1)
		ch := make(chan int)
		go func(ch <-chan int) {
			key := <-ch
			fmt.Println("i =", key)
			go_sync.Done()
		} (ch)
		ch <- i
	}

	go_sync.Wait()
}
