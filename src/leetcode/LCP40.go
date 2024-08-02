package leetcode

import "sort"

//「力扣挑战赛」心算项目的挑战比赛中，要求选手从 `N` 张卡牌中选出 `cnt` 张卡牌，若这 `cnt` 张卡牌数字总和为偶数，则选手成绩「有效」且得分为
// `cnt` 张卡牌数字总和。
//给定数组 `cards` 和 `cnt`，其中 `cards[i]` 表示第 `i` 张卡牌上的数字。 请帮参赛选手计算最大的有效得分。若不存在获取有效得分
//的卡牌方案，则返回 0。
//
//**示例 1：**
//
//> 输入：`cards = [1,2,8,9], cnt = 3`
//>
//> 输出：`18`
//>
//> 解释：选择数字为 1、8、9 的这三张卡牌，此时可获得最大的有效得分 1+8+9=18。
//
//**示例 2：**
//
//> 输入：`cards = [3,3,1], cnt = 1`
//>
//> 输出：`0`
//>
//> 解释：不存在获取有效得分的卡牌方案。
//
//**提示：**
//- `1 <= cnt <= cards.length <= 10^5`
//- `1 <= cards[i] <= 1000`
//
// Related Topics 贪心 数组 排序 👍 114 👎 0

func maxmiumScore(cards []int, cnt int) int {
	// 1.先排序
	sort.Ints(cards)

	// 2.先将后cnt个加起来
	var result int
	var ming, mino = 1000, 1000
	for i := 1; i <= cnt; i++ {
		d := cards[len(cards)-1-(cnt-i)]
		if d%2 == 0 {
			mino = min(mino, d)
		} else {
			ming = min(ming, d)
		}

		result += d
	}

	if result%2 == 0 {
		return result
	}

	// 3.如果不满足，从前找个最大的奇数和自己里面最小的偶数交换；或者找个最大的偶数和自己最小的奇数交换；取两个中最大的
	var maxg, maxo int
	for i := 0; i < len(cards)-cnt; i++ {
		if cards[i]%2 == 0 {
			maxo = max(maxo, cards[i])
		} else {
			maxg = max(maxg, cards[i])
		}
	}

	var result2 int
	if ming != 1000 && maxo > 0 {
		result2 = result - ming + maxo
	}

	if mino != 1000 && maxg > 0 {
		result2 = max(result2, result-mino+maxg)
	}

	return result2
}

//runtime:264 ms
//memory:7.9 MB
