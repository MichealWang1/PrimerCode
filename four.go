package main

import "fmt"

func text1(text map[string]string) {
	text["66"] = " 東莞 "
	fmt.Println(" ++++++++++++++++++++++++++++++++++++++ ")
	for key, value := range text {
		fmt.Println(" key = ", key, " value = ", value)
	}
}

func changeMap(text map[string]string) {
	delete(text, "55")
}

func main() {
	cityMap := make(map[string]string)
	// 添加
	cityMap["11"] = " 北京 "
	cityMap["22"] = " 上海 "
	cityMap["33"] = " 深圳 "
	cityMap["44"] = " 廣州 "
	cityMap["55"] = " 杭州 "
	for key, value := range cityMap {
		fmt.Println(" key = ", key, " value = ", value)
	}
	// 刪除
	delete(cityMap, "11")
	fmt.Println(" --------------------------------------------- ")
	for key, value := range cityMap {
		fmt.Println(" key = ", key, " value = ", value)
	}

	// 修改
	cityMap["44"] = " 广州 "
	fmt.Println(" --------------------------------------------- ")
	for key, value := range cityMap {
		fmt.Println(" key = ", key, " value = ", value)
	}
	// 調用
	text1(cityMap)

	fmt.Println(" --------------------------------------------- ")
	for key, value := range cityMap {
		fmt.Println(" key = ", key, " value = ", value)
	}

	changeMap(cityMap)
	fmt.Println(" --------------------------------------------- ")
	for key, value := range cityMap {
		fmt.Println(" key = ", key, " value = ", value)
	}

}
