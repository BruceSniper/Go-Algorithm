/**
 * @ Author: darkknight
 * @ GitHub: https://github.com/BruceSniper
 * @ Description: 两数之和
 * @ Date: 2023/7/19 01:21
 * @ File: main.go
 **/

package main

import "fmt"

func main() {
	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)
}

func twoSum(nums []int, target int) (result []int) {
	hashTable := map[int]int{}
	for k, v := range nums {
		hashTable[v] = k
	}

	for k, v := range nums {
		if _, ok := hashTable[target-v]; ok {
			if k == hashTable[target-v] {
				continue
			}
			result = []int{k, hashTable[target-v]}
			return
		}
	}

	return
}
