package main

import (
	"fmt"
	"runtime"
)

func main() {
	core := runtime.GOMAXPROCS(1)
	fmt.Println("CPU core =", core)

	go fmt.Println("go")
	runtime.Gosched()
	fmt.Println("golang")
}
