package Leetcode

func maxArea(height []int) int {
	var max int = 0
	i := 0
	j := len(height) - 1
	for i < j {
		if height[i] < height[j] {
			area := height[i] * (j - i)
			if max < area {
				max = area
			}
			i++
		} else {
			area := height[j] * (j - i)
			if max < area {
				max = area
			}
			j--
		}
	}
	return max
}
