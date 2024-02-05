package linkedlist

// partition 分隔链表
func partition(head *ListNode, x int) *ListNode {
	// 模拟，直接拆成两个链表
	// 使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
	// small 比 x 小的
	// large >= x 大的

	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large

	for head != nil {
		// val < x
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			// val >= x 包含其本身
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}

	large.Next = nil
	small.Next = largeHead.Next

	return smallHead.Next
}
