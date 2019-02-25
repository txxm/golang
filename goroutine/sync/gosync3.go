package main

import (
	"fmt"
	"sync"
)

var str string
var mutex sync.Mutex

func main() {
	mutex.Lock()
	go func() {
		str = "goroutine"
		mutex.Unlock()
	}()

	mutex.Lock()
	fmt.Println(str)
	mutex.Unlock()
}
