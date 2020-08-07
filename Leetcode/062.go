package Leetcode

/*
DP:
a 重复性（分治）
    每次只能向下或者向右
    a[i-1][j] + a[i][j-1]
b 定义状态数组
c DP方式
dp[i][j] = dp[i-1][j] + a[i][j-1]
动态规划和递归或者分治没有本质的区别，关键自安于有无
最优的子结构
共性： 找到重复子问题
差异性：最优子结构，中途可以淘汰次优解
*/

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 || m > 100 || n > 100 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
