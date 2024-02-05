package linkedlist

// reverseBetween 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head == nil {
		return nil
	}

	// 不用	反转
	if left == right {
		return head
	}

	// 穿针引线法
	//  left = 2, right = 4
	// A -> B -> C -> D -> E
	// 1. 先找到left 的前一个节点 prev 永远保持不变
	// 2. 当i = left 时， cur = B, next = C
	// 3. 把 c 插入到B 的前面 (即在prev 后面)
	// 4. 把 d 插入到c 的前面

	// 哨兵节点
	dummy := &ListNode{}
	dummy.Next = head

	prev := dummy

	// after loop finished, prev 在left 前面
	// 不能用i < left -2; 因为可能 left == 1
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	cur := prev.Next

	// 遍历次数到right 前一位
	// 几次遍历之后，其实cur 永远指向B
	for i := 0; i < right-left; i++ {
		next := cur.Next

		// 穿针引线
		cur.Next = next.Next  // B -> D
		next.Next = prev.Next // 永远插入到prev 的下一位，C -> B
		prev.Next = next      // A -> C
	}

	return dummy.Next
}
