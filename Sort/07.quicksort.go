package main

import "fmt"

//3, 9, 2, 8, 1, 7, 4, 6, 5, 10

//分3段

//2,1,   3,   9, 8, 7, 4, 6, 5, 10

//1,2,3,   9, 8, 7, 4, 6, 5, 10

//1,2,3,   8, 7, 4, 6, 5, 9, 10

func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(QuickSort(arr))
}

func QuickSort(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		splitdata := arr[0]          //以第一个为基准
		low := make([]int, 0, 0)     //存储比我小的
		high := make([]int, 0, 0)    //存储比我大的
		mid := make([]int, 0, 0)     //存储与我相等的
		mid = append(mid, splitdata) //加入第一个相等的
		for i := 1; i < length; i++ {
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high) //切割递归处理
		myarray := append(append(low, mid...), high...)
		return myarray
	}
}

//func QuickSort2(arr []int, left, right int) []int {
//	length := len(arr) //数组长度
//	if length <= 1 {
//		return arr //一个元素的数组，直接返回
//	} else {
//		splitdata := arr[left]          //以第一个为基准
//		i := left + 1
//		j := right
//		low := make([]int, 0, 0)     //存储比我小的
//		high := make([]int, 0, 0)    //存储比我大的
//		mid := make([]int, 0, 0)     //存储与我相等的
//		mid = append(mid, splitdata) //加入第一个相等的
//		for ; i < length; i++ {
//			for ; j > 0; j-- {
//				if i <= j && arr[i] <= splitdata {
//					i++
//				}
//				if i <= j && arr[j] >= splitdata {
//					j--
//				}
//				if i >= j {
//					break
//				}
//				arr[i], arr[j] = arr[j], arr[i]
//			}
//		}
//		arr[left] = arr[j]
//		arr[j] = splitdata
//		myarray := append(append(low, mid...), high...)
//		return myarray
//	}
//}
