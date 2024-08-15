package leetcode

import "math"

//给你一个由 正整数 组成、大小为 m x n 的矩阵 grid。你可以从矩阵中的任一单元格移动到另一个位于正下方或正右侧的任意单元格（不必相邻）。从值为
//c1 的单元格移动到值为 c2 的单元格的得分为 c2 - c1 。
//
// 你可以从 任一 单元格开始，并且必须至少移动一次。
//
// 返回你能得到的 最大 总得分。
//
//
//
// 示例 1：
//
//
// 输入：grid = [[9,5,7,3],[8,9,6,1],[6,7,14,3],[2,5,3,1]]
//
//
// 输出：9
//
// 解释：从单元格 (0, 1) 开始，并执行以下移动： - 从单元格 (0, 1) 移动到 (2, 1)，得分为 7 - 5 = 2 。 - 从单元格 (2
//, 1) 移动到 (2, 2)，得分为 14 - 7 = 7 。 总得分为 2 + 7 = 9 。
//
// 示例 2：
//
//
//
//
// 输入：grid = [[4,3,2],[3,2,1]]
//
//
// 输出：-1
//
// 解释：从单元格 (0, 0) 开始，执行一次移动：从 (0, 0) 到 (0, 1) 。得分为 3 - 4 = -1 。
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 2 <= m, n <= 1000
// 4 <= m * n <= 10⁵
// 1 <= grid[i][j] <= 10⁵
//
//
// Related Topics 数组 动态规划 矩阵 👍 49 👎 0

// dp[i][j]表示前i行，j列的最大分数
// 最大分数只与起点、终点有关，所以dp[i][j]等于grid[i][j]-左上方子矩阵最小值
// 子矩阵最小值利用动态规划
func maxScore(grid [][]int) int {
	// dp中存放包含自己所在矩阵的最小值
	dp := make([][]int, len(grid))
	for i := range dp {
		dp[i] = make([]int, len(grid[0]))
	}

	result := math.MinInt
	for i, row := range grid {
		for j, cell := range row {
			if i == 0 && j == 0 {
				dp[i][j] = cell
			} else if i == 0 {
				dp[i][j] = min(dp[i][j-1], cell)
				result = max(result, cell-dp[i][j-1])
			} else if j == 0 {
				dp[i][j] = min(dp[i-1][j], cell)
				result = max(result, cell-dp[i-1][j])
			} else {
				t := min(dp[i-1][j], dp[i][j-1])
				dp[i][j] = min(t, cell)
				result = max(result, cell-t)
			}
		}
	}

	return result
}

//runtime:123 ms
//memory:8.4 MB
