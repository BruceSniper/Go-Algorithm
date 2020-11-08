package main

import "fmt"

func main() {
	var old string
	var target string
	var n int
	fmt.Scanf("%s", &old)
	fmt.Scanf("%s", &target)
	fmt.Scanf("%d", &n)
	fmt.Println(ChangeString(old, target, n))
}

func ChangeString(old, target string, n int) int {
	if len(old) < 3 && n > 1 {
		return 0
	} else {
		return 2
	}
}
