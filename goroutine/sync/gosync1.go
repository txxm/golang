package main

import (
	"fmt"
)

var str string

/* 可能不打印任何值，因为print在goroutine之前执行 */
func main() {
	go func() {
		str = "Golang"
	}()
	fmt.Println(str)
}
