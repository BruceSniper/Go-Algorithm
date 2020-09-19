package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //从标准输入生成读对象
	str, _ := reader.ReadString('\n')   //读到换行
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println(StrTest(str))
}

func StrTest(str string) string {
	var result string
	lastIndex := make(map[uint8]int)
	for i := 0; i < len(str); i++ {
		lastIndex[str[i]] = i
	}
	used := make(map[uint8]bool)
	for i := 0; i < len(str); i++ {
		if used[str[i]] {
			continue
		}
		for len(result) != 0 && result[len(result)-1] > str[i] && lastIndex[result[len(result)-1]] > i {
			used[result[len(result)-1]] = false
			result = result[:len(result)-1]
		}
		used[str[i]] = true
		result += string(str[i])
	}
	return result
}
