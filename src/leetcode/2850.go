package leetcode

import (
	"fmt"
	"math"
)

//给你一个大小为 3 * 3 ，下标从 0 开始的二维整数矩阵 grid ，分别表示每一个格子里石头的数目。网格图中总共恰好有 9 个石头，一个格子里可能会有
// 多个 石头。
//
// 每一次操作中，你可以将一个石头从它当前所在格子移动到一个至少有一条公共边的相邻格子。
//
// 请你返回每个格子恰好有一个石头的 最少移动次数 。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[1,1,0],[1,1,1],[1,2,1]]
//输出：3
//解释：让每个格子都有一个石头的一个操作序列为：
//1 - 将一个石头从格子 (2,1) 移动到 (2,2) 。
//2 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
//3 - 将一个石头从格子 (1,2) 移动到 (0,2) 。
//总共需要 3 次操作让每个格子都有一个石头。
//让每个格子都有一个石头的最少操作次数为 3 。
//
//
// 示例 2：
//
//
//
//
//输入：grid = [[1,3,0],[1,0,0],[1,0,3]]
//输出：4
//解释：让每个格子都有一个石头的一个操作序列为：
//1 - 将一个石头从格子 (0,1) 移动到 (0,2) 。
//2 - 将一个石头从格子 (0,1) 移动到 (1,1) 。
//3 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
//4 - 将一个石头从格子 (2,2) 移动到 (2,1) 。
//总共需要 4 次操作让每个格子都有一个石头。
//让每个格子都有一个石头的最少操作次数为 4 。
//
//
//
//
// 提示：
//
//
// grid.length == grid[i].length == 3
// 0 <= grid[i][j] <= 9
// grid 中元素之和为 9 。
//
//
// Related Topics 广度优先搜索 数组 动态规划 矩阵 👍 62 👎 0

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func minimumMoves(grid [][]int) int {
	var more, less [][]int

	// 将大于1块的放入more中，0块的放入less中，核心是吧more中得分配给less的最小值
	for i, row := range grid {
		for j, cell := range row {
			if cell > 1 {
				for k := 0; k < cell-1; k++ {
					more = append(more, []int{i, j})
				}
			} else if cell == 0 {
				less = append(less, []int{i, j})
			}
		}
	}

	// 排列组合，计算每种情况所需移动的步数，找出最小的一个
	res := math.MaxInt32
	// 排列组合的所有情况
	arr := permuteUnique2(more)
	for i := 0; i < len(arr); i++ {
		var t int
		for j := 0; j < len(arr[i]); j++ {
			t += abs(arr[i][j][0]-less[j][0]) + abs(arr[i][j][1]-less[j][1])
		}

		if t < res {
			res = t
		}
	}

	return res
}

// 使用回溯法
func permuteUnique2(nums [][]int) [][][]int {
	res := make([][][]int, 0)
	backTrace2(0, nums, &res)

	return res
}

func backTrace2(index int, nums [][]int, res *[][][]int) {
	if index == len(nums) {
		t := make([][]int, len(nums))
		for i := 0; i < len(nums); i++ {
			t[i] = make([]int, len(nums[i]))
			copy(t[i], nums[i])
		}

		*res = append(*res, t)
		return
	}

	t := make(map[string]bool)
	for i := index; i < len(nums); i++ {
		s := fmt.Sprintf("%d-%d", nums[i][0], nums[i][1])
		if _, ok := t[s]; ok {
			continue
		} else {
			t[s] = true

			swap2(nums, i, index)
			backTrace2(index+1, nums, res)
			swap2(nums, index, i)
		}
	}
}

func swap2(nums [][]int, i, j int) {
	if i != j {
		nums[i][0], nums[i][1], nums[j][0], nums[j][1] = nums[j][0], nums[j][1], nums[i][0], nums[i][1]
	}
}

//runtime:5 ms
//memory:4.6 MB
