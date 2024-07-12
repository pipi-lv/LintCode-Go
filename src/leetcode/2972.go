package leetcode

//给你一个下标从 0 开始的 正 整数数组 nums 。
//
// 如果 nums 的一个子数组满足：移除这个子数组后剩余元素 严格递增 ，那么我们称这个子数组为 移除递增 子数组。比方说，[5, 3, 4, 6, 7]
//中的 [3, 4] 是一个移除递增子数组，因为移除该子数组后，[5, 3, 4, 6, 7] 变为 [5, 6, 7] ，是严格递增的。
//
// 请你返回 nums 中 移除递增 子数组的总数目。
//
// 注意 ，剩余元素为空的数组也视为是递增的。
//
// 子数组 指的是一个数组中一段连续的元素序列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3,4]
//输出：10
//解释：10 个移除递增子数组分别为：[1], [2], [3], [4], [1,2], [2,3], [3,4], [1,2,3], [2,3,4] 和
//[1,2,3,4]。移除任意一个子数组后，剩余元素都是递增的。注意，空数组不是移除递增子数组。
//
//
// 示例 2：
//
//
//输入：nums = [6,5,7,8]
//输出：7
//解释：7 个移除递增子数组分别为：[5], [6], [5,7], [6,5], [5,7,8], [6,5,7] 和 [6,5,7,8] 。
//nums 中只有这 7 个移除递增子数组。
//
//
// 示例 3：
//
//
//输入：nums = [8,7,6,6]
//输出：3
//解释：3 个移除递增子数组分别为：[8,7,6], [7,6,6] 和 [8,7,6,6] 。注意 [8,7] 不是移除递增子数组因为移除 [8,7] 后
//nums 变为 [6,6] ，它不是严格递增的。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁹
//
//
// Related Topics 数组 双指针 二分查找 👍 40 👎 0

func incremovableSubarrayCount2(nums []int) int64 {
	n := len(nums)

	var (
		i            = 0
		j            = n - 1
		result int64 = 1
	)

	last := -1
	// 先遍历头指针到递增序列最后一位，这中间每有一个就是一种
	for k := 0; k < n; k++ {
		if last == -1 || last < nums[k] {
			last = nums[k]
			i = k
			result++
		} else {
			break
		}
	}

	// 如果直接指向最后，说明整个数组都是递增的，直接根据等差数列求和公式计算
	if i == j {
		return int64(n * (n + 1) / 2)
	}

	// 头指针指向最后以后，开始动尾指针，
	// 对于一般情况把右指针从后扫，当 nums[r]<=nums[l] 时，此时把左指针 l 向前移动直到满足 nums[r]>nums[l] 即可。
	// 此时可以删除子数组 [l,r−1]，[l−1,r−1] ... [0,r−1]，一共可贡献 l+2 个答案。计算完 r 的答案后，继续向前移动右指针，当不满足 nums[r]<nums[r+1] 时，停止遍历。
	for j == n-1 || nums[j] < nums[j+1] {
		if i >= 0 {
			if nums[i] < nums[j] {
				j--
				result += int64(i + 2)
			} else {
				// 向前移动i
				i--
			}
		} else {
			j--
			result += 1
		}
	}

	return result
}

//runtime:100 ms
//memory:8.3 MB