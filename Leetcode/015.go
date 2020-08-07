package Leetcode

import "sort"

//a + b + c == 0
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	result := make([][]int, 0)

	//枚举a
	for first := 0; first < n-2; first++ {
		//需要和上一次枚举的值不同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		//枚举c，在a右边的最右边，即队尾
		third := n - 1
		target := -1 * nums[first]
		//枚举b，在a的右边第一个开始，向右移动
		for second := first + 1; second < n-1; second++ {
			//需要和上一次枚举值不同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			//需要保证b指针在c指针的左边
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return result
}

/*
func threeSum(nums []int) [][]int {
    n := len(nums)
    sort.Ints(nums)
    ans := make([][]int, 0)

    // 枚举 a
    for first := 0; first < n; first++ {
        // 需要和上一次枚举的数不相同
        if first > 0 && nums[first] == nums[first - 1] {
            continue
        }
        // c 对应的指针初始指向数组的最右端
        third := n - 1
        target := -1 * nums[first]
        // 枚举 b
        for second := first + 1; second < n; second++ {
            // 需要和上一次枚举的数不相同
            if second > first + 1 && nums[second] == nums[second - 1] {
                continue
            }
            // 需要保证 b 的指针在 c 的指针的左侧
            for second < third && nums[second] + nums[third] > target {
                third--
            }
            // 如果指针重合，随着 b 后续的增加
            // 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
            if second == third {
                break
            }
            if nums[second] + nums[third] == target {
                ans = append(ans, []int{nums[first], nums[second], nums[third]})
            }
        }
    }
    return ans
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/3sum/solution/san-shu-zhi-he-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
