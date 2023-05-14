package main

import (
	"fmt"
	"reflect"
)

type MyBook struct {
	name  string
	auth  string
	title string
	page  int
}

type BookList struct {
	name string
	page int
}

func reflctNum(this interface{}) {
	fmt.Println(" type : ", reflect.TypeOf(this))
	fmt.Println(" value ", reflect.ValueOf(this))

}

// 打印 具体类型 struct里面的成员 type 和 value
func PrintFiledAndMethod(this interface{}) {
	initType := reflect.TypeOf(this)
	fmt.Println(" 1111 具体类型名称 initType = ", initType.Name())

	initValue := reflect.ValueOf(this)
	fmt.Println(" 1111 具体类型名称 initValue = ", initValue)

	fmt.Println(" ======================================= initType.NumField() = ", initType.NumField())
	for i := 0; i < initType.NumField(); i++ {
		childFiled := initType.Field(i)
		childFiledName := childFiled.Name
		childFileType := childFiled.Type
		childFileValue := initValue.Field(i)
		// initValue.Field(i).Interface() 这里报错
		fmt.Println(" 1111 childFiledName  ", childFiledName, " -- childFileType = ", childFileType, " -- childFileValue = ", childFileValue)
	}
}

func main() {
	aa := 14
	// 每个变量都有 一个 pair 对应其 type类型 和  value值
	// pair<type: int, value: 14>
	reflctNum(aa)
	fmt.Println(" ------------------------------ ")
	// pair<type: BookList, value: {"凡人修仙传", 120000}>
	bb := BookList{"凡人修仙传", 120000}
	reflctNum(bb)
	fmt.Println(" ------------------------------ ")
	cc := MyBook{" 紫气东来 ", " 剑舞殇 ", " 穿越修仙 ", 1280}
	PrintFiledAndMethod(cc)
}
