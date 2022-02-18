package main

import (
	"container/list"
)

//给出一棵二叉树，返回其节点值的层次遍历（逐层从左往右访问）

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root: A Tree
 * @return: Level order a list of lists of integer
 */
func levelOrder(root *TreeNode) [][]int {
	// 结果集
	result := make([][]int, 0)

	//层级和树节点对应关系
	fMap := make(map[int]int)

	if root == nil {
		return result
	}

	queue := list.New()
	// 当前节点入队
	queue.PushBack(root)
	fMap[root.Val] = 0

	for queue.Len() > 0 {
		front := queue.Front()
		fNode := front.Value.(*TreeNode)

		if fNode.Left != nil {
			queue.PushBack(fNode.Left)
			fMap[fNode.Left.Val] = fMap[fNode.Val] + 1
		}
		if fNode.Right != nil {
			queue.PushBack(fNode.Right)
			fMap[fNode.Right.Val] = fMap[fNode.Val] + 1
		}

		// 打印，存数组等处理]
		floor := fMap[fNode.Val]
		if floor >= len(result) {
			result = append(result, nil)
		}
		result[floor] = append(result[floor], fNode.Val)

		// 移除元素
		queue.Remove(front)

	}

	return result
}
