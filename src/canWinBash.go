package main

//你正在和朋友玩一个游戏：桌子上有一堆石头，每一次你们都会从中拿出1到3个石头。拿走最后一个石头的人赢得游戏。游戏开始时，你是先手。
//假设两个人都绝对理性，都会做出最优决策。给定石头的数量，判断你是否会赢得比赛。
//举例：有四个石头，那么你永远不会赢得游戏。不管拿几个，最后一个石头一定会被你的朋友拿走。
//输入：n = 4
//输出：False
//解析：先手取走1,2或者3，对方都会取走最后一个
//样例 2：
//
//输入：n = 5
//输出：True
//解析：先手拿1个，必胜

/**
 * @param n: an integer
 * @return: whether you can win the game given the number of stones in the heap
 */
func canWinBash(n int) bool {
	return n % 4 != 0
}
