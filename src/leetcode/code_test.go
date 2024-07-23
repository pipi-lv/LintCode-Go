package leetcode

import "testing"

func TestAccountsMerge(t *testing.T) {
	t.Log(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "john00@mail.com", "johnsmith@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}))
}

func TestMinimumTime(t *testing.T) {
	t.Log(minimumTime(6, [][]int{{0, 4, 7}, {3, 3, 10}, {4, 4, 10}, {1, 0, 1}, {5, 5, 1}, {1, 4, 6}, {3, 5, 9}, {4, 1, 6}, {2, 2, 7}, {3, 4, 2}}, []int{20, 10, 5, 7, 16, 20}))
}

func TestMaximumSum(t *testing.T) {
	t.Log(maximumSum([]int{100, 30, 1, 987, 400, 200, 9}))
}
