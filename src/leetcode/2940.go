package leetcode

// 给你一个下标从 0 开始的正整数数组 heights ，其中 heights[i] 表示第 i 栋建筑的高度。
// 如果一个人在建筑 i ，且存在 i < j 的建筑 j 满足 heights[i] < heights[j] ，那么这个人可以移动到建筑 j 。
// 给你另外一个数组 queries ，其中 queries[i] = [ai, bi] 。第 i 个查询中，Alice 在建筑 ai ，Bob 在建筑 bi 。
// 请你能返回一个数组 ans ，其中 ans[i] 是第 i 个查询中，Alice 和 Bob 可以相遇的 最左边的建筑 。如果对于查询 i ，Alice 和 Bob 不能相遇，令 ans[i] 为 -1 。

// 暴力求解，部分步骤重复计算，超时
func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	res := make([]int, len(queries))

	for i, query := range queries {
		a, b := query[0], query[1]

		if a == b {
			res[i] = a
		} else {
			res[i] = -1
			for j := max(a, b); j < len(heights); j++ {
				if (a == j || (a < j && heights[a] < heights[j])) && (b == j || (b < j && heights[b] < heights[j])) {
					res[i] = j
					break
				}
			}
		}
	}

	return res
}

// ==================== 方法二，利用线段树 ====================

var zd []int

func build(l, r, rt int, heights []int) {
	if l == r {
		zd[rt] = heights[l-1]
		return
	}

	mid := (r + l) >> 1
	build(l, mid, rt<<1, heights)
	build(mid+1, r, rt<<1|1, heights)
	zd[rt] = max(zd[rt<<1], zd[rt<<1|1]) // 由求和改为求最大值
}

// 二分查找
func query(pos, val, l, r, rt int) int {
	if val >= zd[rt] {
		return 0
	}

	if l == r {
		return l
	}

	mid := (l + r) >> 1
	if pos <= mid {
		res := query(pos, val, l, mid, rt<<1)
		if res != 0 {
			return res
		}
	}

	return query(pos, val, mid+1, r, rt<<1|1)
}

func leftmostBuildingQueries2(heights []int, queries [][]int) []int {
	n := len(heights)
	zd = make([]int, n*4)
	build(1, n, 1, heights)

	m := len(queries)
	result := make([]int, m)
	for i := 0; i < m; i++ {
		a, b := queries[i][0], queries[i][1]
		// 交换顺序不影响结果
		if a > b {
			a, b = b, a
		}

		if a == b || heights[a] < heights[b] {
			result[i] = b
			continue
		}

		// 这个时候heights[a] > heights[b],所以从b+1位置开始找第一个大于heights[a]的元素,也肯定大于heights[b],这两个都可以到那个地方
		result[i] = query(b+1, heights[a], 1, n, 1) - 1
	}

	return result
}
