package linkedlist

// mergeTwoLists 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	// 哨兵节点
	dummy := &ListNode{}
	cur := dummy

	// 类似归并排序
	for list1 != nil && list2 != nil {

		val1, val2 := list1.Val, list2.Val

		if val1 <= val2 {
			cur.Next = list1
			list1 = list1.Next

		} else {
			cur.Next = list2
			list2 = list2.Next
		}

		cur = cur.Next

	}

	if list1 != nil {
		cur.Next = list1
	}

	if list2 != nil {
		cur.Next = list2
	}

	return dummy.Next
}
