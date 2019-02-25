package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var go_sync sync.WaitGroup

func show() {
	fmt.Println("goroutine")
}

/* sync.Once()保证*.Do(f)中f()只被调用一次 */
func main() {
	go_sync.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			once.Do(show)
			go_sync.Done()
		}()
	}

	go_sync.Wait()
}
