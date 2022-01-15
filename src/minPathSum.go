package main

func minPathSum(grid [][]int) int {
	cols := len(grid)
	rows := len(grid[0])

	// 初始化结果集
	res := make([][]int, cols)
	for i := 0; i < cols; i++ {
		res[i] = make([]int, rows)
	}

	//dp[i][j] = dp[i-1][j] > dp[i][j-1]?dp[i][j-1]:dp[i-1][j]
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if i == 0 && j == 0 {
				res[0][0] = grid[0][0]
			} else if i == 0 {
				res[i][j] = grid[i][j] + res[i][j-1]
			} else if j == 0 {
				res[i][j] = grid[i][j] + res[i-1][j]
			} else {
				if res[i-1][j] > res[i][j-1] {
					res[i][j] = grid[i][j] + res[i][j-1]
				} else {
					res[i][j] = grid[i][j] + res[i-1][j]
				}
			}
		}
	}

	return res[cols-1][rows-1]
}
