package main

import (
	"fmt"
)

func main() {
	val := 1
	f := function(val)
	fmt.Println("f() =", f())
}

/* 返回值：必包函数。函数返回后，必包状态不会消失,i逃逸到堆上 */
func function(i int) (func() int) {
	return func() int {
		i++
		return i
	}
}
