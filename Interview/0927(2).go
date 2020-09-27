package main

import "fmt"

func searchRange(nums []int, target int) (int, int) {
	left := -1
	right := -1

	if len(nums) == 0 || nums[len(nums)-1] < target {
		return left, right
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			left = i
		}
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] == target {
			right = j
		}
	}
	return right, left
}

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	intArr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &intArr[i])
	}
	fmt.Println(searchRange(intArr, n))
}
