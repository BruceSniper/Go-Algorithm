package Leetcode

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, 0)
	result = append(result, temp)
	recurseSubSets(nums, &temp, 0, &result)
	return result
}

func recurseSubSets(nums []int, tmp *[]int, index int, result *[][]int) {
	if index == len(nums) {
		return
	}

	//1.保存当前值
	*tmp = append(*tmp, nums[index])
	item := make([]int, len(*tmp))
	copy(item, *tmp)
	*result = append(*result, item)
	recurseSubSets(nums, tmp, index+1, result)
	*tmp = (*tmp)[:len(*tmp)-1]

	//2.不保存当前值
	recurseSubSets(nums, tmp, index+1, result)
}
