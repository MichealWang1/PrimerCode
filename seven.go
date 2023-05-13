package main

import "fmt"

type man struct {
	name string
	age  int
}

func (this *man) SetName(str string) {
	this.name = str
}
func (this *man) SetAge(num int) {
	this.age = num
}

type supman struct {
	// 继承
	man
	supmantype string
	level      int
}

func (this *supman) SetSupmantype(str string) {
	this.supmantype = str
}
func (this *supman) Setlevel(level int) {
	this.level = level
}
func main() {
	// 继承初始化1
	t1 := supman{man{"shenzhen", 10}, "superS", 20}
	fmt.Println(" 打印 t1 = ", t1)
	// 继承初始化2
	var t2 supman
	t2.name = " AI周杰伦 "
	t2.age = 20
	t2.supmantype = "superSS"
	t2.level = 40
	fmt.Println(" 打印 t2 = ", t2)

	t1.SetName(" AI孙燕姿 ")
	fmt.Println(" 打印 t1 = ", t1)
}
