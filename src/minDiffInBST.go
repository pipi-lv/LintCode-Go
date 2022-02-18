package main

import "math"

//给定一个二叉搜索树的根结点 root, 返回树中任意两节点的差的最小值。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root:  a Binary Search Tree (BST) with the root node
 * @return: the minimum difference
 */
func minDiffInBST(root *TreeNode) int {
	var result float64

	// 先序遍历获得数组
	var list []int
	preOrder(root, &list)

	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if result == 0 {
				result = math.Abs(float64(list[i] - list[j]))
			} else {
				result = math.Min(result, math.Abs(float64(list[i]-list[j])))
			}
		}
	}

	return int(result)
}

func minDiffInBST2(root *TreeNode) int {
	result, pre := math.MaxInt16, -1

	// 中序遍历获得数组
	var list []int
	midOrder(root, &list)

	for i := 0; i < len(list); i++ {
		if pre != -1 {
			result = min(list[i]-pre, result)
		}
		pre = list[i]
	}

	return result
}

// 二叉搜索树中根遍历是有序的
func midOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	midOrder(root.Left, list)
	*list = append(*list, root.Val)
	midOrder(root.Right, list)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
