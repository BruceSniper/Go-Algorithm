package main

import "fmt"

//9  2  1  6  0  7

//2  9  1  6  0  7		从奇数位开始

//2  1  9  0  6  7		从偶数位开始

//1  2  0  9  6  7		奇

//1  0  2  6  9  7		偶

//0  1  2  6  7  9

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(OddEven(arr))
}

func OddEven(arr []int) []int {

	isSorted := false //奇数位，偶数位都不需要排序的有序

	for isSorted == false {
		isSorted = true
		for i := 1; i < len(arr)-1; i = i + 2 { //奇数位
			if arr[i] > arr[i+1] { //需要交换就换
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}

		}
		for i := 0; i < len(arr)-1; i = i + 2 { //偶数位
			if arr[i] > arr[i+1] { //需要交换就换
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}
