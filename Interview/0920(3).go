package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	Map := make(map[int]int, n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d,%d", &a, &b)
		Map[a] = b
	}
	//fmt.Println(Map)
	fmt.Println(Weight(Map))
}

func Weight(Map map[int]int) int {
	var count int
	for _, v := range Map {
		count += v
	}

	for i := 0; i < count; i++ {

	}
	return count
}
