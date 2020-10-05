package Leetcode

func intersection(nums1 []int, nums2 []int) []int {
	tmpMap := make(map[int]bool, 0)
	result := make([]int, 0)
	Max := findMax(nums1, nums2)
	for i := 0; i < len(Max); i++ {
		tmpMap[Max[i]] = true
	}
	Min := findMin(nums1, nums2)
	for j := 0; j < len(Min); j++ {
		if tmpMap[Min[j]] && findNum(result, Min[j]) {
			result = append(result, Min[j])
		}

	}
	return result
}

func findMax(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return nums1
	}
	return nums2
}

func findMin(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		return nums1
	}
	return nums2
}

func findNum(nums []int, num int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			return false
		}
	}
	return true
}
