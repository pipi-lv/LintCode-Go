package leetcode

//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
// Related Topics 数组 回溯 👍 1583 👎 0

func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	backTraceUnique(0, nums, &res)

	return res
}

func backTraceUnique(index int, nums []int, res *[][]int) {
	if index == len(nums) {
		t := make([]int, len(nums))
		copy(t, nums)
		*res = append(*res, t)
		return
	}

	// 利用map存放状态
	t := make(map[int]bool)
	for i := index; i < len(nums); i++ {
		if _, ok := t[nums[i]]; ok {
			continue
		} else {
			t[nums[i]] = true

			swapUnique(nums, i, index)
			backTraceUnique(index+1, nums, res)
			swapUnique(nums, index, i)
		}
	}
}

func swapUnique(nums []int, i, j int) {
	if i != j {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

//runtime:5 ms
//memory:3.6 MB
