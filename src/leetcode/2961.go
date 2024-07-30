package leetcode

// 给你一个下标从 0 开始的二维数组 variables ，其中 variables[i] = [ai, bi, ci, mi]，以及一个整数 target
// 。
//
// 如果满足以下公式，则下标 i 是 好下标：
//
// 0 <= i < variables.length
// ((aibi % 10)ci) % mi == target
//
// 返回一个由 好下标 组成的数组，顺序不限 。
//
// 示例 1：
//
// 输入：variables = [[2,3,3,10],[3,3,3,1],[6,1,1,4]], target = 2
// 输出：[0,2]
// 解释：对于 variables 数组中的每个下标 i ：
// 1) 对于下标 0 ，variables[0] = [2,3,3,10] ，(2³ % 10)³ % 10 = 2 。
// 2) 对于下标 1 ，variables[1] = [3,3,3,1] ，(3³ % 10)³ % 1 = 0 。
// 3) 对于下标 2 ，variables[2] = [6,1,1,4] ，(6¹ % 10)¹ % 4 = 2 。
// 因此，返回 [0,2] 作为答案。
//
// 示例 2：
//
// 输入：variables = [[39,3,1000,1000]], target = 17
// 输出：[]
// 解释：对于 variables 数组中的每个下标 i ：
// 1) 对于下标 0 ，variables[0] = [39,3,1000,1000] ，(39³ % 10)¹⁰⁰⁰ % 1000 = 1 。
// 因此，返回 [] 作为答案。
//
// 提示：
//
// 1 <= variables.length <= 100
// variables[i] == [ai, bi, ci, mi]
// 1 <= ai, bi, ci, mi <= 10³
// 0 <= target <= 10³
//
// Related Topics 数组 数学 模拟 👍 27 👎 0
// 参见No.50题目，如何快速计算Pow
func getGoodIndices(variables [][]int, target int) []int {
	var result []int

	for i, variable := range variables {
		a, b, c, m := variable[0], variable[1], variable[2], variable[3]

		if powMod(powMod(a, b, 10), c, m) == target {
			result = append(result, i)
		}
	}

	return result
}

func powMod(x int, n int, mod int) int {
	res := 1

	for n > 0 {
		if n&1 == 1 {
			res = res * x % mod
		}

		x = x * x % mod
		n >>= 1
	}

	return res
}
