package main

import "fmt"

//func main() {
//	var m, n int
//	fmt.Scanf("%d", &m)
//	fmt.Scanf("%d", &n)
//	matrix := make([][]int, max(m, n))
//	strArr := make([]string, m)
//	strTmp := make([]string, n)
//
//	for i := 0; i < m; i++ {
//		reader := bufio.NewReader(os.Stdin)    //从标准输入生成读对象
//		strArr[i], _ = reader.ReadString('\n') //读到换行
//		strArr[i] = strings.Replace(strArr[i], "\n", "", -1)
//		strTmp = strings.Split(strArr[i], " ")
//		matrix[i] = make([]int, n)
//		for j := 0; j < n; j++ {
//			matrix[i][j], _ = strconv.Atoi(strTmp[j])
//		}
//	}
//
//	fmt.Println(TestArr(matrix))
//
//}
//
//func max(x, y int) int {
//	if x < y {
//		return y
//	}
//	return x
//}
//
//func TestArr(matrix [][]int) int {
//	if len(matrix) == 0 {
//		return 0
//	}
//	row, col := len(matrix), len(matrix[0])
//	var count int
//
//	for x := 0; x < row; x++ {
//		for y := 0; y < col; y++ {
//			if matrix[x][y] == 1 {
//				count++
//				dfs(x, y, matrix)
//			}
//		}
//	}
//	return count
//}
//
//func dfs(x, y int, matrix [][]int) {
//	if x < 0 || y < 0 || x >= len(matrix) || y >= len(matrix[0]) || matrix[x][y] == 0 {
//		return
//	}
//	matrix[x][y] = 0
//	dfs(x-1,y,matrix)
//	dfs(x+1,y,matrix)
//	dfs(x,y-1,matrix)
//	dfs(x,y+1,matrix)
//}
const N = 110

var m, n int
var g [N][N]int
var st [N][N]bool
var dx, dy = [4]int{-1, 0, 1, 0}, [4]int{0, 1, 0, -1}

func main() {
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &g[i][j])
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func dfs(x, y int) bool {
	if g[x][y] == 0 || st[x][y] {
		return false
	}
	st[x][y] = true
	for i := 0; i < 4; i++ {
		a, b := x+dx[i], y+dy[i]
		if a >= 0 && a < m && b >= 0 && b < n {
			dfs(a, b)
		}
	}
	return true
}
