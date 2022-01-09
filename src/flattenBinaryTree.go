package main

//给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
//展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
//展开后的单链表应该与二叉树 先序遍历 顺序相同。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * @param root: a TreeNode, the root of the binary tree
 * @return: nothing
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	if root.Left == nil {
		return
	}
	flatten(root.Left)
	// 找到左子树的最右子树
	p := root.Left
	for p.Right != nil {
		p = p.Right
	}

	p.Right = root.Right
	root.Right = root.Left
	root.Left = nil
}

/**
 * @param root: a TreeNode, the root of the binary tree
 * 非递归
 * @return: nothing
 */
func flatten2(root *TreeNode) {
	var list []int
	preOrder(root, &list)
	cur := root
	for i := 1; i < len(list); i++ {
		cur.Left = nil
		cur.Right = &TreeNode{Val: list[i], Left: nil, Right: nil}
		cur = cur.Right
	}
}

func preOrder(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	preOrder(root.Left, list)
	preOrder(root.Right, list)
}
