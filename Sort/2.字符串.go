package main

import "fmt"

//func main() {
//	fmt.Println(3 > 2)
//	fmt.Println("12" > "2")
//	fmt.Println("a" < "b") //字符串存在地址
//	//相等0，不等-1，+1
//	//a<b<c 首先比较第一个字母，左边小于右边-1，左边大于右边+1，第一个字母比较不成功标胶第二个，以此类推
//	fmt.Println(strings.Compare("a","b"))
//}

func main() {
	fmt.Println("a" > "b")
	fmt.Println("a" < "b")
	fmt.Println("a" == "a")
	pb := "b"
	pa := "a"
	fmt.Println(&pa, &pb)
	fmt.Println(pa < pb)
}
