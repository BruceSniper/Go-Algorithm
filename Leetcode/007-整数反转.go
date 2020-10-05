package Leetcode

import (
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	str := strconv.Itoa(x)
	newStr := strings.Split(str, "")
	var symbol bool = false
	if newStr[0] == "-" {
		symbol = true
		newStr = newStr[1:]
	}
	length := len(newStr)
	tempStr := make([]string, 0)
	for i := length - 1; i >= 0; i-- {
		tempStr = append(tempStr, newStr[i])
	}
	resultStr := strings.Join(tempStr, "")
	result, _ := strconv.Atoi(resultStr)
	if symbol {
		result = -1 * result
	}
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
}
