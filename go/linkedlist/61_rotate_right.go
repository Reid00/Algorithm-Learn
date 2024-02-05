package linkedlist

// rotateRight 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	// k 有可能大于节点数目，要对其取余
	// 先收尾相连，组成环，然后遍历k次，断开

	if head == nil || k == 0 || head.Next == nil {
		return head
	}

	tail := head

	n := 1

	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	// NOTE 重要: 旋转k次之后，新链表的最后一个节点为
	k = n - k%n

	// 刚好k 是n 的整数倍，也不需要旋转
	if k == n {
		return head
	}

	// 首位相连
	tail.Next = head

	for k > 0 {
		tail = tail.Next
		k--
	}
	// 循环结束后，tail 为新链表的尾部
	newHead := tail.Next
	// 断开
	tail.Next = nil
	return newHead
}
