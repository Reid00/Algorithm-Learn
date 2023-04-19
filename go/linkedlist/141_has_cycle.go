package linkedlist

// https://leetcode.cn/problems/linked-list-cycle/solution/yi-wen-gao-ding-chang-jian-de-lian-biao-wen-ti-h-2/
// 相关总结
// hasCycle 链表是否有环, 双指针 之快慢指针
func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != fast {
		// 如果没有环， fast 会碰到nil
		if fast == nil || fast.Next == nil {
			return false
		}

		// 如果有环，fast 走两步，slow 走一步，在环里早晚会碰撞
		fast = fast.Next.Next
		slow = slow.Next
	}

	return true
}
