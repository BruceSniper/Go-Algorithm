package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * M包糖果，抛M次硬币，硬币连续n次为正面，最多能得到多少颗糖果
 * @param candies int整型一维数组 每包糖果的数量
 * @param coin int整型一维数组 抛硬币的结果
 * @param n int整型 连续是正面的次数
 * @return int整型
 */
func maxCandies(candies []int, coin []int, n int) int {
	// write code here

	M := len(candies)
	countArr := make([]int, 0)
	var MaxInt int
	for i := 0; i < M-n+1; i++ {
		for j := 0; j < n; j++ {
			temp := coin
			temp[i+j] = 0
			tempInt := 0

			for k := 0; k < M; k++ {
				if temp[k] == 0 {
					tempInt += candies[k]
					countArr[k] = tempInt
				}
			}

		}
		for e := 0; e < M; e++ {
			if countArr[e] > MaxInt {
				MaxInt = countArr[e]
			}
		}

	}
	return MaxInt
}

func main() {
	candies := make([]int, 0)
	coin := make([]int, 0)
	var n int
	fmt.Scanf("%v,%v,%d", &candies, &coin, &n)
	fmt.Println(maxCandies(candies, coin, n))
}
