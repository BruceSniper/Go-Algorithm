package Leetcode

/*
算法思路：从空集开始，每增加一个元素，都会与现有各子集组合，
形成新的子集，通过增量迭代的方式实现穷举
*/

func subsets(nums []int) [][]int {
	arr := make([][]int, 0)
	arr = append(arr, []int{})
	for i := 0; i < len(nums); i++ {
		tempArr := make([][]int, 0)
		for _, c := range arr {
			temp := make([]int, 0)
			temp = append(temp, c...)
			temp = append(temp, nums[i]) //增加一个元素
			tempArr = append(tempArr, temp)
		}
		for _, c := range tempArr {
			arr = append(arr, c) //与现有各子集组合， 形成新的子集
		}
	}
	return arr
}
