package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//func Test(str1, str2 string, len1, len2 int) string {
//	if len1 > len2 {
//
//	}
//}

func main() {
	length1 := 0
	fmt.Scanln(&length1)

	reader := bufio.NewReader(os.Stdin) //从标准输入生成读对象
	str1, _ := reader.ReadString('\n')  //读到换行
	str1 = strings.Replace(str1, " ", "", -1)

	length2 := 0
	fmt.Scanln(&length2)
	str2, _ := reader.ReadString('\n') //读到换行
	//str2 = strings.Replace(str2, " ", "", -1)
	fmt.Println(length1)
	fmt.Println(str1)
	fmt.Println(length2)
	fmt.Println(str2)
	//data1, _ := strconv.Atoi(str1) //654321
	//data2, _ := strconv.Atoi(str2)
	//fmt.Println(data1)
	//fmt.Println(data2)
	var data []int
	for length1 > 0 {
		data1, _ := strconv.Atoi(str1[0:1])
		data = append(data, data1)
		str1 = str1[1:]
		length1--
	}
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}

	fmt.Printf("%T", data)
}
