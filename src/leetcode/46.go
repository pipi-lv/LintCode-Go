package leetcode

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
// 示例 2：
//
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//
// 示例 3：
//
//
//输入：nums = [1]
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// nums 中的所有整数 互不相同
//
//
// Related Topics 数组 回溯 👍 2927 👎 0

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	backTrace(0, nums, &res)

	return res
}

func backTrace(index int, nums []int, res *[][]int) {
	// 到最终节点放入结果集
	if index == len(nums) {
		t := make([]int, len(nums))
		copy(t, nums)
		*res = append(*res, t)
		return
	}

	// 向右边的每个节点交换
	for i := index; i < len(nums); i++ {
		swap(nums, i, index)
		backTrace(index+1, nums, res)
		// 回溯之前的状态
		swap(nums, index, i)
	}
}

func swap(nums []int, i, j int) {
	if i != j {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

//runtime:2 ms
//memory:2.7 MB
