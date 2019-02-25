package main

import (
	"fmt"
)

var str string
var ch = make(chan int, 10)

/* 通道完成同步，打印str值 */
func main() {
	go func() {
		str = "Golang"
		ch <- 0
	}()

	<-ch /* 只有ch发送完成才能执行接收 */
	fmt.Println(str)
}
