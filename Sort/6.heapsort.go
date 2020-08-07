package main

import "fmt"

//堆排序

func HeapSortMax(arr []int, length int) []int {
	//length := len(arr)
	if length <= 1 {
		return arr
	} else {
		depth := length/2 - 1         //深度，n 2*n+1,2*n+2
		for i := depth; i >= 0; i-- { //循环所有的三节点
			topmax := i //假定最大的在i的位置
			leftchild := 2*i + 1
			rightchild := 2*i + 2
			if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越界
				topmax = leftchild //如果左边比顶点大，纪录最大
			}
			if rightchild <= length-1 && arr[rightchild] > arr[topmax] { //防止越界
				topmax = rightchild //如果右边比顶点大，纪录最大
			}
			if topmax != i { //确保i的值就是最大
				arr[i], arr[topmax] = arr[topmax], arr[i]
			}
		}

	}
	return arr
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastmesslen := length - i //每次截取一段
		HeapSortMax(arr, lastmesslen)
		fmt.Println(arr)
		isneedexchange := false
		if i < length {
			arr[0], arr[lastmesslen-1] = arr[lastmesslen-1], arr[0]
			isneedexchange = true
		}
		if !isneedexchange {
			break
		}
		fmt.Println("ex", arr)

	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}

	//fmt.Println(HeapSortMax(arr, 8))
	fmt.Println(HeapSort(arr))
}
