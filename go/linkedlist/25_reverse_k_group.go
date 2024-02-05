package linkedlist

// reverseKGroup K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {

	// 哨兵
	dummy := &ListNode{}
	dummy.Next = head

	cur := head
	prev := dummy

	for cur != nil {
		tail := prev

		// 循环k 次 找到 k区间的tail
		for i := 0; i < k; i++ {
			tail = tail.Next
			// tail 为nil 说明 剩余元素 不够k个，不用反转
			if tail == nil {
				return dummy.Next
			}
		}

		// 下 k 个元素的开始
		next := tail.Next

		newHead, newTail := reverse(cur, tail)

		// 前k个 后k个 和这段链表关联起来
		prev.Next = newHead
		newTail.Next = next

		// 更新下k 个元素的prev
		prev = newTail
		cur = next
	}

	return dummy.Next

}

// reverse 反转指定区间的链表
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	// k = 2
	// A -> B -> C -> D -> E
	// B -> A -> C -> D -> E  第一次 反转之后A 指向C， A 是cur， 所以c 是pre

	prev := tail.Next

	cur := head

	// 循环终止的条件是prev 到tail 这个位置
	for prev != tail {
		next := cur.Next

		// 反转
		cur.Next = prev
		prev = cur
		cur = next
	}

	return tail, head
}
