package Leetcode

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]]++
	}

	for _, v := range numMap {
		if v > 1 {
			return true
		}
	}
	return false
}
