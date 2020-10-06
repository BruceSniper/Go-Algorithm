package Leetcode

func reverseArr(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

func rotate(nums []int, k int) {
	reverseArr(nums)
	reverseArr(nums[:k%len(nums)])
	reverseArr(nums[k%len(nums):])
}
