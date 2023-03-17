package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}

	var head *ListNode

	var carry int

	for l1 != nil || l2 != nil || carry != 0 {
		var add1, add2 int
		if l1 != nil {
			add1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			add2 = l2.Val
			l2 = l2.Next
		}

		sum := add1 + add2 + carry

		val := sum % 10
		carry = sum / 10

		if head == nil {
			head = &ListNode{
				Val: val,
			}
			dummy.Next = head
		} else {
			node := &ListNode{
				Val: val,
			}
			head.Next = node
			head = node
		}
	}
	return dummy.Next
}
