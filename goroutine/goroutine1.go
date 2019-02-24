package main

import (
	"fmt"
	"sync"
)

func main() {
	var go_sync sync.WaitGroup
	i := 6

	go_sync.Add(1)
	go function(i, &go_sync)
	go_sync.Wait()
}

func function(i int, go_sync *sync.WaitGroup) {
	fmt.Println("i =", i)
	go_sync.Done()
}
