package main

import (
	"fmt"
)

func isValidTest(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}

	Map := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	Stack := []byte{}
	for j := 0; j < length; j++ {
		if Map[s[j]] > 0 {
			if len(Stack) == 0 || Stack[len(Stack)-1] != Map[s[j]] {
				return false
			}
			Stack = Stack[:len(Stack)-1]
		} else {
			Stack = append(Stack, s[j])
		}
	}
	return len(Stack) == 0
}

func main() {
	strArr := make([]string, 100)
	for i := 0; i < 100; i++ {
		fmt.Scanf("%s", &strArr[i])
		if strArr[i] == "" {
			break
		}
	}
	for i := 0; i < 100; i++ {
		if strArr[i] == "" {
			break
		}
		fmt.Println(isValidTest(strArr[i]))

	}
}
