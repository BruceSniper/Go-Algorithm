package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str1 string
	var str3 string
	fmt.Scanln(&str1)                   //原单词
	reader := bufio.NewReader(os.Stdin) //从标准输入生成读对象
	str2, _ := reader.ReadString('\n')  //读到换行
	str2 = strings.Replace(str2, "\n", "", 1)
	fmt.Scanln(&str3) //替换的新单词
	fmt.Println(Replace(str1, str2, str3))
}

//func Replace(str1 , str2, str3 string) string {
//	String := make([]string,0)
//	String = strings.Split(str2, " ")
//	for i := 0; i < len(String)-1; i++ {
//		if String[i] == str1 || checkInclusion(str1, String[i]){
//			strRepalce(String, i, str3)
//		}
//	}
//	return strings.Join(String, " ")
//}

func strRepalce(strArr []string, i int, str string) {
	strArr[i] = str
}

func checkInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range s1 {
		cnt1[s1[i]-'a']++
		cnt2[s2[i]-'a']++
	}
	for i := 0; i < len(s2)-len(s1); i++ {
		if cnt1 == cnt2 {
			return true
		}
		cnt2[s2[i]-'a']--
		cnt2[s2[i+len(s1)]-'a']++
	}
	return cnt1 == cnt2
}

func Replace(str1, str2, str3 string) string {
	String1 := strings.Split(str2, ",")
	var resultString string
	var tempStr1 []string
	var tempStr2 []string
	for i := 0; i < len(String1); i++ {
		tempStr1[i] = String1[i]
		tempStr2 = strings.Split(tempStr1[i], " ")
	}
	for i := 0; i < len(tempStr2[i]); i++ {
		if tempStr2[i] == str1 || checkInclusion(str1, tempStr2[i]) {
			strRepalce(tempStr2, i, str3)
		}
	}
	return strings.Join(String, " ")
}
