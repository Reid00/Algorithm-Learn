package linkedlist

// reverseList 迭代版本
func reverseList(head *ListNode) *ListNode {

	var prev *ListNode

	cur := head

	for cur != nil {
		// 只处理cur
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}

// reverseList2 递归版本
func reverseList2(head *ListNode) *ListNode {

	// 中止条件
	if head == nil || head.Next == nil {
		return head
	}

	// 递归反复调用 反转 head.Next 之后的链表
	newHead := reverseList2(head.Next)

	// 处理 head 和 head.Next 的反转问题
	head.Next.Next = head
	head.Next = nil

	return newHead

}
