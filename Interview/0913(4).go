package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	fmt.Println(Test(n, m))
}

func Test(n, m int) int {
	var count int
	var start int
	for i := 0; i < 3*n; i++ {
		start = m
		count += start + m*i
	}

	return count
}
