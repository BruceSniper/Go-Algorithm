package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //从标准输入生成读对象
	str1, _ := reader.ReadString('\n')  //读到换行
	str1 = strings.Replace(str1, "\n", "", -1)
	result := IsYear(str1)
	length := len(result)
	for i := 0; i < length; i++ {
		fmt.Print(result[i])
		fmt.Print(" ")
	}
}

func IsYear(str string) []string {
	find := make([]string, len(str))
	strArr := make([]string, len(str))
	if len(str) <= 2000 {
		strArr = strings.Split(str, " ")
		regx, _ := regexp.Compile(`[1-3][0-9]{1,3}`)
		for i := 0; i < len(strArr); i++ {
			if len(strArr[i]) == 4 {
				find = regx.FindAllString(str, -1)
			}
		}

	}
	return find
}

//import java.util.Scanner;
//import java.util.regex.Matcher;
//import java.util.regex.Pattern;
//
//class Main {
//public static void main(String[] args){
//Scanner scanner = new Scanner(System.in);
//String reg="[1-3][0-9][0-9][0-9]";
//String str= scanner.nextLine();
//Pattern pattern = Pattern.compile(reg);
//Matcher matcher = pattern.matcher(str);
//while(matcher.find()){
//System.out.print(matcher.group()+" ");
//}
//}
//}
