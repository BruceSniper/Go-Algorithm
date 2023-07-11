/**
 * @ Author: darkknight
 * @ GitHub: https://github.com/BruceSniper
 * @ Description: 016-最接近的三数之和
 * @ Date: 2023/7/10 23:27
 * @ File: main.go
 **/

package main

import (
	"math"
	"sort"
)

func main() {
	threeSumClosest([]int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2)
}

func threeSumClosest(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < len(nums)-2; i++ {

			if i > 0 && nums[i] == nums[i-1] {
				continue
			}

			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}

	return res
}

func abs(v int) int {
	if v > 0 {
		return v
	} else {
		return -1 * v
	}
}
