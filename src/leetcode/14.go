package leetcode

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
//
// Related Topics 字典树 字符串 👍 3146 👎 0

func longestCommonPrefix(strs []string) string {
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		var j int

		for j < len(result) {
			if j < len(strs[i]) {
				if strs[i][j] == result[j] {
					j++
				} else {
					break
				}
			} else {
				break
			}
		}

		result = result[:j]
	}

	return result
}
