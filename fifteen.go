package main

import "fmt"

var chan2 = make(chan int)
var chan3 = make(chan int)

func SendMsg2() {
	x, y := 1, 1
	for {
		select {
		case chan2 <- x:
			x = y
			y = x + y
		case <-chan3:
			fmt.Println(" 退出函数SendMsg2 ")
			return
		}
	}
}

func main() {
	defer close(chan2)
	defer close(chan3)
	go func() {
		for i := 0; i < 10; i++ {
			num := <-chan2
			fmt.Println("  打印 chan2 取出的值 num = ", num)
		}
		chan3 <- 0
	}()
	SendMsg2()

}
