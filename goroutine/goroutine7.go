package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

/* 在单核中，for循环占据CPU，say没有机会执行,无法打印 */
func main() {
	runtime.GOMAXPROCS(1)
	go say("world")
	for {
	}
}
