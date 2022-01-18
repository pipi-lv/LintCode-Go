package main

//给定一个只含非负整数的m*nm∗n网格，找到一条从左上角到右下角的可以使数字和最小的路径。
//你在同一时间只能向下或者向右移动一步

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
