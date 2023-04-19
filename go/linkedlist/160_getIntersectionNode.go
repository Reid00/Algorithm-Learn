package linkedlist

// getIntersectionNode 相交链表
// headA 在A链表走到底部之后，从B 链表的head 开始走，
// 同理 headB 走完之后，走A， 因为相交之后的长度相同， A+B 的长度相同
// 所以在相交处 会汇合
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB

	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}

	return pA

}
