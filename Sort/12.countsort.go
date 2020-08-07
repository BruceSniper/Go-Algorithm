package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 4, 3, 2, 1, 2, 5, 5, 3, 4, 3, 2, 1}
	fmt.Println(Countsort(arr))
}

func Countsort(arr []int) []int {
	max := SelectSortMaxY(arr)         //寻找最大值
	sortedarr := make([]int, len(arr)) //排序之后存储
	countsarr := make([]int, len(arr)) //统计次数

	for _, v := range arr {
		countsarr[v]++
	}
	fmt.Println("第一次统计次数", countsarr) //统计次数
	for i := 1; i <= max; i++ {
		countsarr[i] += countsarr[i-1] //叠加
	}
	fmt.Println("次数叠加", countsarr) //统计次数
	for _, v := range arr {
		sortedarr[countsarr[v]-1] = v //展开数据

		countsarr[v]-- //递减
		fmt.Println("zkcount", countsarr)
		fmt.Println("zk", sortedarr)
	}
	return sortedarr
}

func SelectSortMaxY(arr []int) int {
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
