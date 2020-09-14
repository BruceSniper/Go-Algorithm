package main

import "fmt"

func firstMissingPositive(nums []int) int {
	// write code here
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] <= 0 {
			nums[i] = length + 1
		}
	}
	for i := 0; i < length; i++ {
		num := Abs(nums[i])
		if num <= length {
			fmt.Println(num - 1)
			nums[num-1] = -Abs(nums[num-1])
		}
	}
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return length + 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
