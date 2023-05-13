package main

import (
	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"

	"fmt"
)

func ExistInt(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func ExistString(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

type K interface {
	int | string | float32
}

type Node[T interface{ int | string | float32 }] struct {
	Data T
}

func Exist[T K](s []T, i T) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

type S interface{ ~string }

type name string

func ExistName[Text S](str []Text, i Text) bool {
	for _, v := range str {
		if v == i {
			return true
		}
	}
	return false
}

func Filter[S, T any](arr []S, f func(S) T) []T {
	dest := make([]T, len(arr))
	for i, _ := range arr {
		dest[i] = f(arr[i])
	}
	return dest
}

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println(len(string("你好sz!")))
	fmt.Println(" ExistInt([]int{}, 5) ", ExistInt([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(" ExistString([]string{} c) ", ExistString([]string{"a", "v", "c"}, "c"))

	fmt.Println(" ExistName([]string{} c) ", ExistName([]string{"a", "v", "c"}, "c"))

	//fmt.Println(" ExistString([]int{}, 5) ", ExistString([]int{1, 2, 3, 4, 5}, 5))
	//fmt.Println(" ExistInt([]string{} c) ", ExistInt([]string{"a", "v", "c"}, "c"))

	fmt.Println(" Exist([]int{}, 5) ", Exist([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(" Exist([]string{} c) ", Exist([]string{"a", "v", "c"}, "c"))

	n1 := Node[int]{
		Data: 1,
	}
	fmt.Println(" 打印 n1.Data = ", n1.Data)
	n2 := Node[string]{
		Data: "1111",
	}
	fmt.Println(" 打印 n2.Data = ", n2.Data)

	strudent := []User{
		{Name: "aaaa", Age: 12},
		{Name: "bbbb", Age: 13},
		{Name: "cccc", Age: 11},
		{Name: "dddd", Age: 14},
	}
	strName := Filter(strudent, func(u User) string {
		return u.Name
	})

	strAge := Filter(strudent, func(u User) int {
		return u.Age
	})

	fmt.Println(" 名字 = ", strName)
	fmt.Println(" 年龄 = ", strAge)
}
