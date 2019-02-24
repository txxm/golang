package main

import (
	"fmt"
	"time"
)

/* goroutine启动需要时间，此时循环执行完，i的值是5且必包传引用，打印值相同*/
func main() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("&i = %p, i = %d\n", &i, i)
		}()
	}

	time.Sleep(1 * time.Second)
}
