package main

import (
	"fmt"
	"sync"
)

/* 多个goroutine但是输出结果是无序的 */
func main() {
	var go_sync sync.WaitGroup

	for i := 0; i < 5; i++ {
		go_sync.Add(1)
		go function(i, &go_sync)
	}
	go_sync.Wait()
}

func function(i int, go_sync *sync.WaitGroup) {
	fmt.Println("i =", i)
	go_sync.Done()
}
