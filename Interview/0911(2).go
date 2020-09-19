package main

import (
	"fmt"
)

func main() {
	var len1,len2 int
	fmt.Scanf("%d %d", &len1, &len2) //第一行两个长度n,m

	strMg := make(map[int]int, len2)
	key := 0
	value := 0
	for i:=0; i<len2; i++ {
		fmt.Scanf("%d %d", &strMg[key]value)
	}
	fmt.Println(strMg)