package main

import "fmt"

// 当变量名称用 大写 则 是公有属性 小写则是私有属性
// 结构体中 变量名大写 则是公有属性小写则是私有属性
// 结构体名称是大写开头则 表示其他包也能访问 改结构体
// 函数名大写 也是一样 如果函数名大写则其他包都能访问 小写则只有当前go文件能使用
type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this Hero) Show() {
	fmt.Println("this.Name = ", this.Name)
	fmt.Println(" this.Ad ", this.Ad)
	fmt.Println("this.Level ", this.Level)
}

func (this Hero) SetHeroName(str string) {
	this.Name = str
}
func (this Hero) SetHeroAd(num int) {
	this.Ad = num
}
func (this Hero) SetHeroLevel(num int) {
	this.Level = num
}

func (this *Hero) HeroName(str string) {
	this.Name = str
}
func (this *Hero) HeroAd(num int) {
	this.Ad = num
}
func (this *Hero) HeroLevel(num int) {
	this.Level = num
}
func main() {
	supeMan := Hero{
		Name:  "",
		Ad:    0,
		Level: 0,
	}
	supeMan.Show()
	supeMan.SetHeroName(" 超人1 ")
	supeMan.SetHeroAd(11)
	supeMan.SetHeroLevel(111)
	// 这个时候改变超人没有成功 值传递 概念
	supeMan.Show()
	// 调用指针的 就可以修改
	supeMan.HeroName(" 超人2 ")
	supeMan.HeroAd(22)
	supeMan.HeroLevel(222)
	supeMan.Show()
}
