package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 无参函数 无返回值
	go func() {
		defer fmt.Println(" A.defer ")
		func() {
			defer fmt.Println(" B.defer ")
			fmt.Println(" 函数 B ")
			runtime.Goexit()
		}()
		fmt.Println(" 函数A ")
	}()
	// 有参有返回值
	go func(a int, b int) int {
		c := a + b
		fmt.Println(" 函数c c = ", c)
		return c
	}(10, 20)

	i := 0
	for {
		i++
		if i == 10 {
			break
		}
		fmt.Println(" 主线程 i = ", i)
		time.Sleep(1 * time.Second)
	}
}
