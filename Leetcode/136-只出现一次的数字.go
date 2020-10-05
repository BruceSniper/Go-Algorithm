package Leetcode

//func singleNumber(nums []int) int {
//	numMap := make(map[int]int, 0)
//	for i := 0; i < len(nums); i++ {
//		numMap[nums[i]]++
//	}
//	for k, v := range numMap {
//		if v == 1 {
//			return k
//		}
//	}
//	return 0
//}

func singleNumber(nums []int) int {
	single := 0
	for _, v := range nums {
		single ^= v
	}
	return single
}
