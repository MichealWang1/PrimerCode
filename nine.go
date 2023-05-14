package main

import "fmt"

type BookType struct {
	name string
}

func showType(this interface{}) {
	fmt.Println(" 打印this ", this)
	// 类型断言  必须是 空的 interface
	value, ok := this.(string)
	if ok {
		fmt.Println(" 是字符串 ", value)
	}
	value1, ok1 := this.(int)
	if ok1 {
		fmt.Println(" 是整数类型 ", value1)
	}
	value2, ok2 := this.(BookType)
	if ok2 {
		fmt.Println(" 是BookType ", value2)
	}
}

func main() {
	text := BookType{" 凡人修仙传 "}
	showType(text)
	showType(" 1234 ")
	showType(123)
}
