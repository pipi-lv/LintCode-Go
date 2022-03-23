package main

//删除链表中等于给定值 val 的所有节点。

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @param head: a ListNode
 * @param val: An integer
 * @return: a ListNode
 */
func removeElements(head *ListNode, val int) *ListNode {
	temp := head
	var pre *ListNode

	for temp != nil {
		if temp.Val == val {
			if pre != nil {
				pre.Next = temp.Next
			} else {
				head = temp.Next
			}
		} else {
			pre = temp
		}
		temp = temp.Next
	}
	return head
}
