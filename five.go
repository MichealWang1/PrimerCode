package main

import "fmt"

type myint int32
type Book struct {
	name  string
	auth  string
	title string
	page  int
}

func change(text Book) {
	text.title = " 作者妄語 "

	fmt.Println(" text = ", text)
}
func main() {
	var a myint = 10
	fmt.Println(" a = ", a)
	var likebook Book
	likebook.name = " 凡人修仙傳 "
	likebook.auth = " 韓立韓老魔 "
	likebook.title = " 韓立修仙傳 "
	likebook.page = 1280

	fmt.Println(" likebook = ", likebook)

	b := "a"
	c := "阿"
	fmt.Println(" b = ", b, "len(b) = ", len(b))
	fmt.Println(" c = ", c, "len(c) = ", len(c))
	change(likebook)
	// 得出 結構體也是值拷貝
	fmt.Println(" 修改後 likebook = ", likebook)
}
