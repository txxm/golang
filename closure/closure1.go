package main

import (
	"fmt"
)

/* 必包函数使用外部变量m时，此处为引用，而非传值，所以打印的地址相同 */
func main() {
	var i int = 3
	m := i
	fmt.Printf("&m = %p, m = %d\n", &m, m)

	func() {
		m++
		fmt.Printf("&m = %p, m = %d\n", &m, m)
	}()
	if m == i {
		fmt.Printf("必包函数内为数据副本，m值未变化, m = %d\n", m)
	} else {
		fmt.Printf("必包函数内为数据引用，m值变化, m = %d\n", m)
	}

	function(i)
}


/* 必包函数传参n，传递的参数为n的副本，而非引用，地址会发生变化 */
func function(j int) {
	n := j
	fmt.Println("************************")
	fmt.Printf("&n = %p, n = %d\n", &n, n)

	func(n int) {
		n++
		fmt.Printf("&n = %p, n = %d\n", &n, n)
	}(n)
	if n == j {
		fmt.Printf("必包函数内为数据副本，n值未变化, n = %d\n", n)
	} else {
		fmt.Printf("必包函数内为数据引用，n值变化, n = %d\n", n)
	}
}
