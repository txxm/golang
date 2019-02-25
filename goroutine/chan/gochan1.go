package main

import (
	"fmt"
	"runtime"
)

var quit chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		runtime.Gosched()
		fmt.Printf("%d", i)
	}
	quit <- 0
}

func main() {
	runtime.GOMAXPROCS(2)

	go loop()
	go loop()

	for i := 0; i < 2; i++ {
		<-quit
	}
}
