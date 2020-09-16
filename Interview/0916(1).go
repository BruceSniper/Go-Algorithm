package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	intArr := make([]int, n)
	intArr = fib(n)
	var j int
	for i := 1; i <= n; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print(intArr[j])
			fmt.Print(" ")
		}
		for k := j - 2; k >= 1; k-- {
			fmt.Print(intArr[k])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func fib(n int) []int {
	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	for i := 2; i <= n; i++ {
		res[i] = (res[i-1] + res[i-2]) % 1000000007
	}
	return res
}
