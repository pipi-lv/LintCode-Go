package leetcode

//给你一个下标从 0 开始的字符串 num ，表示一个非负整数。
//
// 在一次操作中，您可以选择 num 的任意一位数字并将其删除。请注意，如果你删除 num 中的所有数字，则 num 变为 0。
//
// 返回最少需要多少次操作可以使 num 变成特殊数字。
//
// 如果整数 x 能被 25 整除，则该整数 x 被认为是特殊数字。
//
//
//
//
//
// 示例 1：
//
//
//输入：num = "2245047"
//输出：2
//解释：删除数字 num[5] 和 num[6] ，得到数字 "22450" ，可以被 25 整除。
//可以证明要使数字变成特殊数字，最少需要删除 2 位数字。
//
// 示例 2：
//
//
//输入：num = "2908305"
//输出：3
//解释：删除 num[3]、num[4] 和 num[6] ，得到数字 "2900" ，可以被 25 整除。
//可以证明要使数字变成特殊数字，最少需要删除 3 位数字。
//
// 示例 3：
//
//
//输入：num = "10"
//输出：1
//解释：删除 num[0] ，得到数字 "0" ，可以被 25 整除。
//可以证明要使数字变成特殊数字，最少需要删除 1 位数字。
//
//
//
//
// 提示
//
//
// 1 <= num.length <= 100
// num 仅由数字 '0' 到 '9' 组成
// num 不含任何前导零
//
//
// Related Topics 贪心 数学 字符串 枚举 👍 54 👎 0

func minimumOperations(num string) int {
	// 25 50 75 00 结尾
	// 只保留0
	// 全部删除为0

	// 找所有2 5 0 7的位置
	// 首先找5 和 0的位置，如果不存在则只能全部删除
	result := len(num)
	for i, n := range num {
		if n == '5' { // 往前找2、7或往后找0
			for j := 0; j < i; j++ {
				if num[j] == '2' || num[j] == '7' {
					result = min(result, (i-j-1)+(len(num)-1-i))
				}
			}

			for j := i + 1; j < len(num); j++ {
				if num[j] == '0' {
					result = min(result, (j-i-1)+(len(num)-1-j))
				}
			}
		} else if n == '0' { // 往前找0或往后找0，不需要再往前找5了，因为上面已经找过了
			for j := 0; j < i; j++ {
				if num[j] == '0' {
					result = min(result, (i-j-1)+(len(num)-1-i))
				}
			}

			for j := i + 1; j < len(num); j++ {
				if num[j] == '0' {
					result = min(result, (j-i-1)+(len(num)-1-j))
				}
			}

			result = min(result, len(num)-1)
		}
	}

	return result
}

//runtime:0 ms
//memory:2.2 MB
