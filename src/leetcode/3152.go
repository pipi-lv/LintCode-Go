package leetcode

// 如果数组的每一对相邻元素都是两个奇偶性不同的数字，则该数组被认为是一个 特殊数组 。
//
// 你有一个整数数组 nums 和一个二维整数矩阵 queries，对于 queries[i] = [fromi, toi]，请你帮助你检查 子数组
// nums[fromi..toi] 是不是一个 特殊数组 。
//
// 返回布尔数组 answer，如果 nums[fromi..toi] 是特殊数组，则 answer[i] 为 true ，否则，answer[i] 为
// false 。
//
// 示例 1：
//
// 输入：nums = [3,4,1,2,6], queries = [[0,4]]
//
// 输出：[false]
//
// 解释：
//
// 子数组是 [3,4,1,2,6]。2 和 6 都是偶数。
//
// 示例 2：
//
// 输入：nums = [4,3,1,6], queries = [[0,2],[2,3]]
//
// 输出：[false,true]
//
// 解释：
//
// 子数组是 [4,3,1]。3 和 1 都是奇数。因此这个查询的答案是 false。
// 子数组是 [1,6]。只有一对：(1,6)，且包含了奇偶性不同的数字。因此这个查询的答案是 true。
//
// 提示：
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
// 1 <= queries.length <= 10⁵
// queries[i].length == 2
// 0 <= queries[i][0] <= queries[i][1] <= nums.length - 1
//
// Related Topics 数组 二分查找 前缀和 👍 36 👎 0

// 暴力求解
func isArraySpecialII_1(nums []int, queries [][]int) []bool {
	var arr []int // 这边和下面主要区别是下面存了特殊数组长度，而这边存了特殊位置坐标
	pre := nums[0]%2 == 0
	for i, num := range nums[1:] {
		t := num%2 == 0
		if pre == t {
			arr = append(arr, i+1)
		}

		pre = t
	}

	result := make([]bool, len(queries))
l:
	for i, query := range queries {
		s, e := query[0], query[1]
		for _, t := range arr {
			if s <= t-1 && e >= t {
				result[i] = false
				continue l
			}
		}

		result[i] = true
	}

	return result
}

//runtime:754 ms
//memory:18 MB

// 方法二：动态规划
// 我们知道对于任意的特殊数组 nums[i⋯j]，则该数组的任意连续子数组也一定也是特殊数组，因此我们只需要检测以 j 为结尾的最长特殊子数组是否可以覆盖区间 [i,j]。
// 假设我们已知以索引 j 为结尾的最长特殊数组的长度 dp[j]，此时只需要判断 dp[j] 是否大于等于区间 [i,j] 的长度即可，即此时是否满足 dp[j]≥j−i+1，如果满足，则此时 nums[i⋯j] 一定为特殊子数组。

// 根据题意可以知，对于每个索引 i 的最长特殊数组的长度 dp[i] 计算方法如下：
// 1.如果 nums[i] 与左边相邻的元素 nums[i−1] 奇偶性相同，则此时 dp[i]=1;
// 2.如果 nums[i] 与左边相邻的元素 nums[i−1] 奇偶性不同，则此时 nums[i] 可以追加到以 nums[i−1] 为结尾的最长特殊数组的后面，则 dp[i]=dp[i−1]+1;
// 在判断两个元素奇偶性是否相同时，可以利用位运算来实现，对于给定的元素 a,b，当满足 (a^b)&1=1 时，则 a,b 的奇偶性不同，否则奇偶性相同；
func isArraySpecialII_2(nums []int, queries [][]int) []bool {
	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 1; i < len(nums); i++ {
		if (nums[i]^nums[i-1])&1 == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}

	result := make([]bool, len(queries))
	for i, query := range queries {
		s, e := query[0], query[1]
		if dp[e] >= e-s+1 {
			result[i] = true
		}
	}

	return result
}
