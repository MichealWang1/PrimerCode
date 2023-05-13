package main

import "fmt"

func text(aa []int) {
	aa = append(aa, 100)
	aa = append(aa, 200)
	aa = append(aa, 300)
	fmt.Println("aa = ", aa)
}

func main() {
	// slicel := make([]int, 3)
	// var slicel []int = make([]int, 3)
	var slicel = []int{1, 2, 3}
	fmt.Println("slicel = ", slicel)
	fmt.Println(" len(slicel) = ", len(slicel), "  cap(slicel) = ", cap(slicel))
	// 追加 1個元素 切片滿了  則  增加容量 等于原容量 * 2
	slicel = append(slicel, 4)

	fmt.Println("slicel = ", slicel)
	fmt.Println(" len(slicel) = ", len(slicel), "  cap(slicel) = ", cap(slicel))
	text(slicel)
	// 切片調用時 是 值傳遞
	fmt.Println(" len(slicel) = ", len(slicel), "  cap(slicel) = ", cap(slicel), " slicel = ", slicel)
	// 切片 截取
	s1 := slicel[0:4]
	fmt.Println(" len(s1) = ", len(s1), "  cap(s1) = ", cap(s1), " s1 = ", s1)
	// 截取時，得到的 s1 的切片 地址 也是指向  slicel 的地址
	s1[0] = 1000
	fmt.Println(" len(slicel) = ", len(slicel), "  cap(slicel) = ", cap(slicel), " slicel = ", slicel)

}
