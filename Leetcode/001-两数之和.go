package Leetcode

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
				break
			}
		}
	}
	return result
}
