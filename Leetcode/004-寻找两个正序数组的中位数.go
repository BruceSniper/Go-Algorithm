package Leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	tmpArr := make([]int, m+n)
	i, j := 0, 0
	var result float64
	for k := 0; k < len(tmpArr); k++ {
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				tmpArr[k] = nums1[i]
				i++
			} else {
				tmpArr[k] = nums2[j]
				j++
			}
		} else if i < m {
			tmpArr[k] = nums1[i]
			i++
		} else if j < n {
			tmpArr[k] = nums2[j]
			j++
		}
	}

	if len(tmpArr)%2 == 0 {
		result = float64(tmpArr[len(tmpArr)/2]+tmpArr[len(tmpArr)/2-1]) / 2
	} else {
		result = float64(tmpArr[len(tmpArr)/2])
	}
	return result
}
