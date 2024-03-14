package divideandconquer

// mergeKLists 合并 K 个升序链表
// METHOD3: 最小堆
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	// 此处忽略
	return nil
}

// METHOD2: 归并
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return SplitList(lists, 0, len(lists))

}

// SplitList 二分
func SplitList(lists []*ListNode, l, r int) *ListNode {

	// 左右索引相同，返回其中一个
	if l == r {
		return lists[l]
	}

	if l > r {
		return nil
	}

	mid := l + (r-l)>>1

	return mergeTwoList(SplitList(lists, l, mid), SplitList(lists, mid+1, r))
}

// METHOD1: 暴力
func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var res *ListNode

	// 暴力，两两合并两个ListNode
	for _, list := range lists {
		res = mergeTwoList(res, list)
	}
	return res
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var dummy = &ListNode{}
	cur := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}

		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}

	return dummy.Next
}
