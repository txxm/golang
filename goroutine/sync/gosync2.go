package main

import (
	"fmt"
)

var str string
var a string
var b string

/* 缓冲通道 */
var ch = make(chan int, 10)
/* 无缓冲通道 */
var channel = make(chan int)

/* 无缓冲通道同时准备好才可收发数据 */
func function() {
	go func() {
		a = "goroutine no buffer channel"
		<-channel
	}()

	channel <- 0
	fmt.Println(a)
}

/* 不保证打印 */
func function1() {
	go func() {
		b = "goroutine."
		<-ch
	}()

	ch <- 0
	fmt.Println(b)
}

/* 通道完成同步，打印str值 */
func main() {
	go func() {
		str = "Golang"
		ch <- 0 /* 或close(ch) */
	}()

	<-ch /* 只有ch发送完成才能执行接收 */
	fmt.Println(str)

	fmt.Println("*************")
	function()

	fmt.Println("*************")
	function1()
}
