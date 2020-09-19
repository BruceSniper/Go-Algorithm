package main

import "fmt"

func generateMatrix(m, n int) [][]int {
	result := make([][]int, m)
	if m > 0 && m < 10000 && n > 0 && n < 10000 {
		for i := 0; i < n; i++ {
			result = append(result, make([]int, n))
		}

		//num := 65
		//nRound := m / 2
		//if nRound*2 != m {
		//	nRound = nRound + 1
		//}
		//for nr := 0; nr < nRound; nr++ {
		//	tempLen := n - 2*nr - 1
		//	if tempLen == 0 {
		//		result[nr][nr] = num
		//		continue
		//	}
		//	if tempLen < 0 {
		//		break
		//	}
		//	for i := 0; i < tempLen; i++ {
		//		result[nr][nr+1] = num
		//		num++
		//	}
		//	for i := 0; i < tempLen; i++ {
		//		result[nr+i][n-nr-1] = num
		//		num++
		//	}
		//	for i := tempLen; i > 0; i-- {
		//		result[n-nr-1][nr+i] = num
		//		num++
		//	}
		//	for i := tempLen; i > 0; i-- {
		//		result[nr+i][nr] = num
		//		num++
		//	}
		//}
		size := m * n
		a, b := 0, n-1 //左右边界
		c, d := 0, m-1 //上下边界
		idx := 65
		for idx <= size {
			//向右走
			for i := a; i <= b; i++ {
				result[c][i] = idx
				idx++
			}
			c++ //上边界加1
			//向下走
			for i := c; i <= d; i++ {
				result[i][b] = idx
				idx++
			}
			b-- //右边界减1
			//向左走
			for i := b; i >= a; i-- {
				result[d][i] = idx
				idx++
			}
			d--
			//向上走
			for i := d; i >= c; i-- {
				result[i][a] = idx
				idx++
			}
			a++
		}
	}
	return result
}

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(generateMatrix(i, j))
			fmt.Print(" ")
		}
		fmt.Println()

	}

}
