package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func majorityElement(nums []int) int {
	length := len(nums)
	sort.Ints(nums)
	return nums[length/2]
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
	fmt.Println(majorityElement(numArr))
}
