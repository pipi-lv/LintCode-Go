package leetcode

//给你两个数组 nums 和 andValues，长度分别为 n 和 m。
//
// 数组的 值 等于该数组的 最后一个 元素。
//
// 你需要将 nums 划分为 m 个 不相交的连续 子数组，对于第 iᵗʰ 个子数组 [li, ri]，子数组元素的按位 AND 运算结果等于
//andValues[i]，换句话说，对所有的 1 <= i <= m，nums[li] & nums[li + 1] & ... & nums[ri] ==
//andValues[i] ，其中 & 表示按位 AND 运算符。
//
// 返回将 nums 划分为 m 个子数组所能得到的可能的 最小 子数组 值 之和。如果无法完成这样的划分，则返回 -1 。
//
//
//
// 示例 1：
//
//
// 输入： nums = [1,4,3,3,2], andValues = [0,3,3,2]
//
//
// 输出： 12
//
// 解释：
//
// 唯一可能的划分方法为：
//
//
// [1,4] 因为 1 & 4 == 0
// [3] 因为单元素子数组的按位 AND 结果就是该元素本身
// [3] 因为单元素子数组的按位 AND 结果就是该元素本身
// [2] 因为单元素子数组的按位 AND 结果就是该元素本身
//
//
// 这些子数组的值之和为 4 + 3 + 3 + 2 = 12
//
// 示例 2：
//
//
// 输入： nums = [2,3,5,7,7,7,5], andValues = [0,7,5]
//
//
// 输出： 17
//
// 解释：
//
// 划分 nums 的三种方式为：
//
//
// [[2,3,5],[7,7,7],[5]] 其中子数组的值之和为 5 + 7 + 5 = 17
// [[2,3,5,7],[7,7],[5]] 其中子数组的值之和为 7 + 7 + 5 = 19
// [[2,3,5,7,7],[7],[5]] 其中子数组的值之和为 7 + 7 + 5 = 19
//
//
// 子数组值之和的最小可能值为 17
//
// 示例 3：
//
//
// 输入： nums = [1,2,3,4], andValues = [2]
//
//
// 输出： -1
//
// 解释：
//
// 整个数组 nums 的按位 AND 结果为 0。由于无法将 nums 划分为单个子数组使得元素的按位 AND 结果为 2，因此返回 -1。
//
//
//
// 提示：
//
//
// 1 <= n == nums.length <= 10⁴
// 1 <= m == andValues.length <= min(n, 10)
// 1 <= nums[i] < 10⁵
// 0 <= andValues[j] < 10⁵
//
//
// Related Topics 位运算 线段树 队列 数组 二分查找 动态规划 👍 36 👎 0

func minimumValueSum(nums []int, andValues []int) int {
	const INF = (1 << 20) - 1 // 20个1的2进制数字，和任何数组&，都等于那个数字
	m, n := len(nums), len(andValues)

	cache := make([]map[int]int, n*m)
	for i := 0; i < len(cache); i++ {
		cache[i] = make(map[int]int)
	}

	// nums的第i个元素，作为andValues的第j组的最后一个元素，第j组的最小值
	var dfs func(i, j, cur int) int
	dfs = func(i, j, cur int) int {
		key := i*n + j
		if i == m && j == n { // 两个同时到头，停止遍历
			return 0
		}

		if i == m || j == n { // 只有一个到头，不作为结果
			return INF
		}

		if v, ok := cache[key][cur]; ok { // 还没有遍历到结尾，看看是否之前遍历过
			return v
		}

		cur &= nums[i]
		if cur&andValues[j] < andValues[j] { // 不需要继续遍历，不可能遍历出来
			return INF
		}

		// 继续遍历下一个和换组遍历两种情况比较找最小值
		res := dfs(i+1, j, cur)
		if cur == andValues[j] {
			res = min(res, dfs(i+1, j+1, INF)+nums[i])
		}

		// 缓存
		cache[key][cur] = res
		return res
	}

	if res := dfs(0, 0, INF); res < INF {
		return res
	}

	return -1
}

//runtime:279 ms
//memory:45.5 MB
