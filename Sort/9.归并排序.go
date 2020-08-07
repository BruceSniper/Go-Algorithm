package main

import "fmt"

//1, 9, 2, 8, 3, 7, 4, 6, 5, 10

//1, 9, 2,   8, 3,    7, 4, 6, 5, 10

//1, 2, 9    3, 8,

//1, 2, 3, 8, 9		//4, 5, 6, 7, 10

//1, 2, 3, 4, 5, 6, 7, 8, 9, 10

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	fmt.Println(MergeSort(arr))
}

func MergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		mid := length / 2
		leftarr := MergeSort(arr[:mid])
		rightarr := MergeSort(arr[mid:])

		return Merge(leftarr, rightarr)
	}

}

func Merge(leftarr []int, rightarr []int) []int {
	leftindex := 0    //左边的索引
	rightindex := 0   //右边的索引
	var lastarr []int //最终的数组
	for leftindex < len(leftarr) && rightindex < len(rightarr) {
		if leftarr[leftindex] < rightarr[rightindex] {
			lastarr = append(lastarr, leftarr[leftindex])
			leftindex++
		} else if leftarr[leftindex] > rightarr[rightindex] {
			lastarr = append(lastarr, rightarr[rightindex])
			rightindex++
		} else {
			leftindex++
			rightindex++
		}
	}
	for leftindex < len(leftarr) { //把没有结束的归并过来
		lastarr = append(lastarr, leftarr[leftindex])
		leftindex++
	}
	for rightindex < len(rightarr) { //把没有结束的归并过来
		lastarr = append(lastarr, rightarr[rightindex])
		rightindex++
	}
	return lastarr
}
