package main

import "fmt"

func main() {
	// 定义一个channel 无缓冲
	c := make(chan int)
	go func(a int, b int) {
		defer fmt.Println(" goroutine A 结束 ")
		fmt.Println(" goroutine A 正在运行 ")
		c <- a + b
	}(10, 20)

	num := <-c
	fmt.Println(" 得到 goroutine A 返回值 num = ", num)

	// 关闭 channal
	close(c)
}
