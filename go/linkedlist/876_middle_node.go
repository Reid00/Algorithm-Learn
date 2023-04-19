package linkedlist

// middleNode 双指针，fast  每次走两步，slow 走一步 fast/fast.Next == nil
// slow 所在位置就是middle Node
// 长度是偶数返回第二个
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
