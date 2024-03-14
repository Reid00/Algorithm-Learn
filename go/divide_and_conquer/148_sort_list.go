package divideandconquer

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// sortList 排序链表
func sortList(head *ListNode) *ListNode {

	// 1. 利用快慢指针将链表拆分，一分为二，递归拆到一边一个
	// 2. merge 每个ListNode
	return SplitListNode(head)
}

func SplitListNode(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head
	// slow 前面的节点
	preSlow := head

	// 快慢指针，fast 走两步，slow 走一步，最后fast 到最后的时候，slow 在中间
	for fast != nil && fast.Next != nil {
		preSlow = slow
		fast = fast.Next.Next
		slow = slow.Next
	}

	// 截断
	preSlow.Next = nil
	l := SplitListNode(head)
	r := SplitListNode(slow)
	// 合并
	return merge(l, r)
}

// merge 合并链表
func merge(l, r *ListNode) *ListNode {
	if l == nil {
		return r
	}

	if r == nil {
		return l
	}

	var dummy = &ListNode{}
	cur := dummy

	for l != nil && r != nil {
		if l.Val < r.Val {
			cur.Next = l
			l = l.Next
		} else {
			cur.Next = r
			r = r.Next
		}
		cur = cur.Next
	}

	// 上面循环结束之后，l,r 必定有一个为nil
	if l != nil {
		cur.Next = l
	}

	if r != nil {
		cur.Next = r
	}

	return dummy.Next
}
