package main

import "fmt"

func main() {
	var mymap map[string]string
	fmt.Println(" len(mymap) = ", len(mymap))
	mymap = make(map[string]string, 10)

	mymap["11"] = "1111"
	mymap["22"] = "2222"
	mymap["33"] = "3333"
	mymap["44"] = "4444"
	mymap["55"] = "5555"
	fmt.Println(" len(mymap) = ", len(mymap), " mymap = ", mymap)

	mymap1 := make(map[int]string)
	mymap1[1] = "1111"
	mymap1[2] = "2222"
	mymap1[3] = "3333"
	mymap1[4] = "4444"
	mymap1[5] = "5555"
	fmt.Println(" len(mymap1) = ", len(mymap1), " mymap1 = ", mymap1)

	mymap2 := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
	}
	fmt.Println(" len(mymap2) = ", len(mymap2), " mymap1 = ", mymap2)

	mymap2["sex"] = "6"

	fmt.Println(" len(mymap2) = ", len(mymap2), " mymap1 = ", mymap2)
}
