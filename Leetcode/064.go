package Leetcode

func minPathSum(grid [][]int) int {
	length := len(grid)
	size := len(grid[0])
	dp := make([][]int, length)
	if length <= 0 || size <= 0 {
		return 0
	}
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, size)
	}
	//初始化
	dp[0][0] = grid[0][0]
	// 初始化最左边的列
	for i := 1; i < length; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// 初始化最上边的行
	for j := 1; j < size; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// 推导出 dp[m-1][n-1]
	for i := 1; i < length; i++ {
		for j := 1; j < size; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
		}
	}

	return dp[length-1][size-1]
}
