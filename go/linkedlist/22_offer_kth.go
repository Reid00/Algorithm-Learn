package linkedlist

// getKthFromEnd 倒数第k 个元素
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {

		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
