package main

import (
	"fmt"
)

func main() {
	var len1, len2 int
	fmt.Scanf("%d %d", &len1, &len2) //第一行两个长度n,m
	var str string
	fmt.Scanln(&str) //第二行长度n，想查询的字符串

	strMg := make([]string, len2)
	for i := 0; i < len2; i++ {
		fmt.Scanf("%s", &strMg[i])
	}
	fmt.Println(strMg)
	//reader := bufio.NewReader(os.Stdin) //从标准输入生成读对象
	//str1, _ := reader.ReadString('\n')  //读到换行
	//str1 = strings.Replace(str1, " ", "", -1)
	//length := int(str1[0])   //字符串长度n
	//lengthMg := int(str1[1]) //敏感词长度m
	//var str string
	//fmt.Scan(&str) //第二行长度n，想查询的字符串
	//strMg := make([]string, 6)
	//for i := 0; i < lengthMg; i++ {
	//	fmt.Scanln(&strMg[i]) //接下来读取m行
	//}
	//fmt.Println(Found(str, length, strMg))
}

//func Found(str string, length int, mg []string) int {
//	var count int
//	var Word = make(map[string]int)
//	for key := range mg {
//		Word[mg[key]] = 0
//	}
//	fmt.Println(Word)
//	for i := 0; i < length-1; i++ {
//		for j := i + 1; j < length; j++ {
//			if _, ok := Word[str[i:j]]; ok {
//				Word[str[i:j]] ++
//			}
//		}
//	}
//
//	for _, v := range Word {
//		count += v
//	}
//
//	return count
//}

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//)
//
//
//func main() {
//	for {
//		fmt.Print(">")
//		reader := bufio.NewReader(os.Stdin)
//		bytes, _, _ := reader.ReadLine()
//		str := string(bytes)
//		fmt.Println(str)
//		}
//	}
//}
