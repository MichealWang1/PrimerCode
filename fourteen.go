package main

import (
	"fmt"
)

// 创建一个 全局的 channal
var c1 = make(chan int, 10)

func SendMsg() {
	// 函数调用 全局 channal 发送数据

	// 关闭 channal
	defer close(c1)

	for i := 0; i < 12; i++ {
		c1 <- i
		fmt.Println(" 子go协城 发送 c1 i = ", i)
	}
}

func main() {
	// 带缓冲 channal
	go SendMsg()
	/*for {
		if num, ok := <-c1; ok {
			fmt.Println(" 1111 channal 取出数据  num = ", num)
		} else {
			break
		}
	}*/
	for value := range c1 {
		fmt.Println(" 2222 channal 取出数据  value = ", value)
	}
	fmt.Println(" 关闭 main() ")

}
