package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println(" newTask 打印i i = ", i)
		time.Sleep(1 * time.Second)
	}
}

// 协程 gomaxprocs
func main() {
	fmt.Println(" ------------------- ")
	go newTask()

	i := 0
	for {
		i++
		fmt.Println(" main 打印i i = ", i)
		time.Sleep(1 * time.Second)
		if i == 10 {
			break
		}
	}
}
