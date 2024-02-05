package linkedlist

// deleteDuplicates 删除排序链表中的重复元素 II （升序排列）
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	// 哨兵
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy

	cur := head

	// 要比较cur 和 next 的val 所以要保证非nil
	for cur != nil {

		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}
		// 如果相同说明 cur 没有发生移动，则无重复值
		if prev.Next == cur {
			prev = cur
		} else {
			prev.Next = cur.Next
		}

		cur = cur.Next
	}

	return dummy.Next
}

// deleteDuplicates 删除排序链表中的重复元素 II （升序排列）
func deleteDuplicates2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	// 哨兵
	dummy := &ListNode{}
	dummy.Next = head

	// 要保证 如果head 是重复的被删掉，所以当前应该在一个确定不为重复的元素，即哑节点
	cur := dummy

	// 要比较cur 和 next 的val 所以要保证非nil
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			rmVal := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == rmVal {
				cur.Next = cur.Next.Next
			}

		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
