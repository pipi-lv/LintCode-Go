package leetcode

import "sort"

//给你两个整数数组 nums1 和 nums2。
//
//从 nums1 中移除两个元素，并且所有其他元素都与变量 x 所表示的整数相加。如果 x 为负数，则表现为元素值的减少。
//
//执行上述操作后，nums1 和 nums2 相等 。当两个数组中包含相同的整数，并且这些整数出现的频次相同时，两个数组 相等 。
//
//返回能够实现数组相等的 最小 整数 x 。
//
//
//
//示例 1:
//
//输入：nums1 = [4,20,16,12,8], nums2 = [14,18,10]
//
//输出：-2
//
//解释：
//
//移除 nums1 中下标为 [0,4] 的两个元素，并且每个元素与 -2 相加后，nums1 变为 [18,14,10] ，与 nums2 相等。
//
//示例 2:
//
//输入：nums1 = [3,5,5,3], nums2 = [7,7]
//
//输出：2
//
//解释：
//
//移除 nums1 中下标为 [0,3] 的两个元素，并且每个元素与 2 相加后，nums1 变为 [7,7] ，与 nums2 相等。
//
//
//
//提示：
//
//3 <= nums1.length <= 200
//nums2.length == nums1.length - 2
//0 <= nums1[i], nums2[i] <= 1000
//测试用例以这样的方式生成：存在一个整数 x，nums1 中的每个元素都与 x 相加后，再移除两个元素，nums1 可以与 nums2 相等。

// 自己写的愚蠢方法，暴力搜索
func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums2)

	result := 1000
	for i := 0; i < len(nums1); i++ {
	l:
		for j := i + 1; j < len(nums1); j++ {
			var t []int
			for k := 0; k < len(nums1); k++ {
				if k == i || k == j {
					continue
				}

				t = append(t, nums1[k])
			}

			o := 0
			pre := -1001
			sort.Ints(t)
			for _, v := range t {
				if pre == -1001 {
					pre = nums2[o] - v
				} else if nums2[o]-v != pre {
					continue l
				}
				o++
			}

			result = min(result, pre)
		}
	}

	return result
}

// 优秀写法，利用第一个节点一定在前三个来反推出所有整个数组，又通过倒过来循环来直接得出最小值，很巧妙
func minimumAddedInteger2(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	// 结果数组[0]一定在排序后的nums1中的前三个
	for i := 2; i >= 0; i-- {
		// 双指针找出剩余结果数组
		left, right := i+1, 1
		for left < len(nums1) && right < len(nums2) {
			if nums1[left]-nums2[right] == nums1[i]-nums2[0] {
				left++
				right++
			} else {
				left++
			}
		}
		if right == len(nums2) {
			return nums2[0] - nums1[i]
		}
	}

	return 0
}
