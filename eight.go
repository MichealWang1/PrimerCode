package main

import (
	"fmt"
)

type Animal interface {
	Sleep()           // 动物睡觉
	GetColor() string // 获取动物颜色
	GetType() string  // 获取动物类型
}
type Cat struct {
	color   string
	strType string
}

func (this *Cat) Sleep() {
	fmt.Println(" 猫类动物要睡觉了 ")
}
func (this *Cat) GetColor() string {
	fmt.Println(" 动物的颜色是 ", this.color)
	return this.color
}
func (this *Cat) GetType() string {
	fmt.Println(" 动物的类型是 ", this.strType)
	return this.strType
}

type Dog struct {
	color   string
	strType string
}

func (this *Dog) Sleep() {
	fmt.Println(" 狗类动物要睡觉了 ")
}
func (this *Dog) GetColor() string {
	fmt.Println(" 动物的颜色 ", this.color)
	return this.color
}
func (this *Dog) GetType() string {
	fmt.Println(" 动物的类型是 ", this.strType)
	return this.strType
}

func showAnimal(this Animal) {
	this.Sleep()
	this.GetColor()
	this.GetType()
}

func main() {
	var animal Animal
	animal = &Cat{" 灰色 ", " 大猫 "}
	animal.Sleep()
	animal.GetColor()
	animal.GetType()
	fmt.Println("----------------------------------------")
	animal = &Dog{" 黄色 ", " 老狗"}
	animal.Sleep()
	animal.GetColor()
	animal.GetType()
	fmt.Println("----------------------------------------")
	// 这种 最常见的 面向对象的方法
	tiger := Cat{" 花色 ", " 老虎 "}
	mouse := Dog{" 白色 ", " 老鼠 "}
	showAnimal(&tiger)
	fmt.Println("----------------------------------------")
	showAnimal(&mouse)
}
