package main

import "fmt"

//1 9 2 8 7 6 4 5

//9 	1 2 8 7 6 4 5

//9  8		1 2 7 6 4 5

//9  8  7  	1 2 6 4 5

//9  8  7  6	1 2 4 5

//9  8  7  6  5		1 2 4

//9  8  7  6  5  4		1 2

//9  8  7  6  5  4  2		1

//9  8  7  6  5  4  2  1

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(SelectSort(arr))
}

func SelectSortMax(arr []int) int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr[0]
	} else {
		max := arr[0] //假定第一个最大
		for i := 1; i < length; i++ {
			if arr[i] > max { //任何一个比我大的数，最大的
				max = arr[i]
			}
		}
		return max
	}

}

/*
	选择排序
*/

func SelectSort(arr []int) []int {
	length := len(arr) //数组的长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		for i := 0; i < length-1; i++ { //只剩一个元素，不需要挑选
			min := i                          //标记索引
			for j := i + 1; j < length; j++ { //每次选出一个极小值
				if arr[min] > arr[j] {
					min = j //保存极小值的索引
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] //数据交换
			}
			fmt.Println(arr)
		}
		return arr
	}
}
