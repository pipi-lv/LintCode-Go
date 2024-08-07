package leetcode

//给你一个 二进制字符串 s 和一个整数数组 queries ，其中 queries[i] = [firsti, secondi] 。
//
// 对于第 i 个查询，找到 s 的 最短子字符串 ，它对应的 十进制值 val 与 firsti 按位异或 得到 secondi ，换言之，val ^
//firsti == secondi 。
//
// 第 i 个查询的答案是子字符串 [lefti, righti] 的两个端点（下标从 0 开始），如果不存在这样的子字符串，则答案为 [-1, -1] 。如
//果有多个答案，请你选择 lefti 最小的一个。
//
// 请你返回一个数组 ans ，其中 ans[i] = [lefti, righti] 是第 i 个查询的答案。
//
// 子字符串 是一个字符串中一段连续非空的字符序列。
//
//
//
// 示例 1：
//
//
//输入：s = "101101", queries = [[0,5],[1,2]]
//输出：[[0,2],[2,3]]
//解释：第一个查询，端点为 [0,2] 的子字符串为 "101" ，对应十进制数字 5 ，且 5 ^ 0 = 5 ，所以第一个查询的答案为 [0,2]。第二个
//查询中，端点为 [2,3] 的子字符串为 "11" ，对应十进制数字 3 ，且 3 ^ 1 = 2 。所以第二个查询的答案为 [2,3] 。
//
//
// 示例 2：
//
//
//输入：s = "0101", queries = [[12,8]]
//输出：[[-1,-1]]
//解释：这个例子中，没有符合查询的答案，所以返回 [-1,-1] 。
//
//
// 示例 3：
//
//
//输入：s = "1", queries = [[4,5]]
//输出：[[0,0]]
//解释：这个例子中，端点为 [0,0] 的子字符串对应的十进制值为 1 ，且 1 ^ 4 = 5 。所以答案为 [0,0] 。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s[i] 要么是 '0' ，要么是 '1' 。
// 1 <= queries.length <= 10⁵
// 0 <= firsti, secondi <= 10⁹
//
//
//
//
// Related Topics 位运算 数组 哈希表 字符串 👍 27 👎 0

// substringXorQueries 这道题核心在于查找时间过多，所以需要先穷举所有情况放到哈希表中，这样可以大幅减少查找时间
func substringXorQueries(s string, queries [][]int) [][]int {
	dict := make(map[int][]int)
	for i, ss := range s {
		if ss == '0' {
			dict[0] = []int{i, i}
			break
		}
	}

	for i, c := range s {
		if c == '0' {
			continue
		}

		// 预处理 30 位的原因是因为数据量最大是 10^9，对应二进制最多 30 位
		for r, x := i, 0; r < len(s) && r < i+30; r++ {
			x = x<<1 | int(s[r]&1) // 这步很精妙 ‘0’&1=0 ‘1’&1=1 实现了字符串转数字
			if _, ok := dict[x]; !ok {
				dict[x] = []int{i, r}
			}
		}
	}

	result := make([][]int, len(queries))
	notFound := []int{-1, -1} // 避免重复创建
	for i, q := range queries {
		p, ok := dict[q[0]^q[1]]
		if !ok {
			result[i] = notFound
		} else {
			result[i] = p
		}
	}

	return result
}
