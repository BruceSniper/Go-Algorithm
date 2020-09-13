package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func main() {
	strArr := make([]string, 0)
	reader := bufio.NewReader(os.Stdin) //从标准输入生成读对象
	str1, _ := reader.ReadString('\n')  //读到换行
	str1 = strings.Replace(str1, "\n", "", -1)
	strArr = strings.Split(str1, " ")
	numArr := make([]int, len(strArr))
	for i := 0; i < len(strArr); i++ {
		temp, _ := strconv.Atoi(strArr[i])
		numArr[i] = temp
	}
	result := make([][]int, len(numArr))
	result = threeSum(numArr)
	for i := 0; i < len(numArr)/3; i++ {
		fmt.Printf("%d %d %d\n", result[i][0], result[i][1], result[i][2])
	}
}
