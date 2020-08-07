package main

import "fmt"

//1, 9, 2, 8, 3, 7, 4, 6, 5, 10

//1				 7

//   9				4

//		2			   6

//		   8			 5

//			 3				10

//1	 4	2  5 3	7  9   6 8  10

//1			 3			 8

//   4			7			10

//		2		   9

//		   5		   6

//1		   5	  9			10

//	 4		 3		   6

//		2		7		 8

//1  3  2  5 4  7  9   6  8  10

//1     2    4     9      8

//   3     5    7      6     10

//1     2    4     8      9

//   3     5    6      7     10

//1  3  2  5 4  6  8   7  9  10

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 4, 6, 5, 10}
	//fmt.Println(SelectSortMax(arr))
	fmt.Println(ShellSort(arr))
}

func ShellSort(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {   //一个元素的数组，直接返回
		return arr
	} else {
		gap := length / 2
		for gap > 0 {
			for i := 0; i < gap; i++ { //处理每个元素的步长
				ShellSortStep(arr, i, gap)
			}
			gap /= 2 //gap--
		}
	}
	return arr
}

func ShellSortStep(arr []int, start, gap int) {
	length := len(arr)
	for i := start + gap; i < length; i += gap { //插入排序的变种
		backup := arr[i] //备份插入的数据
		j := i - gap     //上一个位置循环找到位置插入
		for j >= 0 && backup < arr[j] {
			arr[j+gap] = arr[j] //从前往后移动
			j -= gap
		}
		arr[j+gap] = backup //插入
		fmt.Println(arr)

	}
}
