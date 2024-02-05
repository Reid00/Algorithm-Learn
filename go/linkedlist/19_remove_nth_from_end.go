package linkedlist

// removeNthFromEnd 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// 遍历两遍，第一遍先求长度
	if head == nil {
		return head
	}

	// get Length
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	// 删除倒数第n 个
	dummy := &ListNode{}
	dummy.Next = head
	cur := head

	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	// 删除元素，把 Next 指针 指向下一个元素
	cur.Next = cur.Next.Next

	return dummy.Next
}
