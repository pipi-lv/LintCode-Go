package leetcode

//
//
// 给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。如果存在，返回 true ；否则
//，返回 false 。
//
//
//
// 二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
//
//
//
// 示例 1：
//
//
//输入：root = [3,4,5,1,2], subRoot = [4,1,2]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
//输出：false
//
//
//
//
// 提示：
//
//
// root 树上的节点数量范围是 [1, 2000]
// subRoot 树上的节点数量范围是 [1, 1000]
// -10⁴ <= root.val <= 10⁴
// -10⁴ <= subRoot.val <= 10⁴
//
//
// Related Topics 树 深度优先搜索 二叉树 字符串匹配 哈希函数 👍 1071 👎 0

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if subRoot == nil || root == nil {
		return false
	}

	if root.Val == subRoot.Val {
		// TODO：这边的优化方向为：当前值相同时，先判断左右子树高度和sub高度是否相等，相等再去判断是否是子树
		if isSame(root, subRoot) {
			return true
		}
	}

	if isSubtree(root.Left, subRoot) {
		return true
	}

	return isSubtree(root.Right, subRoot)
}

// 自己的丑陋写法
func isSame(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if subRoot == nil || root == nil {
		return false
	} else {
		if root.Val == subRoot.Val {
			if !isSame(root.Left, subRoot.Left) {
				return false
			}

			return isSame(root.Right, subRoot.Right)
		}
	}

	return false
}

// 别人的高级写法
func isSame2(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}

	return root.Val == subRoot.Val && isSame(root.Left, subRoot.Left) && isSame(root.Right, subRoot.Right)
}
