package linkedlist

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// copyRandomList 随机链表的复制
// O(1) 空间复杂度
func copyRandomList(head *Node) *Node {
	// 因为要深拷贝
	// 例如对于链表 A→B→C, 第一遍 复制A' 表示复制节点，作为原节点的后驱，建立如下的关系
	// 我们可以将其拆分为 A→A′→B→B′→C→C′
	// 后续再遍历 建立random 关系
	// 再遍历，建立next 关系

	if head == nil {
		return nil
	}

	for node := head; node != nil; node = node.Next.Next {
		// 关系由 A->B   => A -> A' -> B
		nodeNew := &Node{Val: node.Val, Next: node.Next}
		node.Next = nodeNew
	}

	// 组织random 关系
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			// node.Random 是原始节点的random 值，node.Randmo.Next 才是copy 值
			node.Next.Random = node.Random.Next
		}
	}

	// 组织next 关系
	newHead := head.Next
	for node := head; node != nil; node = node.Next {

		nodeCopy := node.Next
		// 恢复原来的链表结构
		node.Next = node.Next.Next

		if nodeCopy.Next != nil {
			nodeCopy.Next = nodeCopy.Next.Next
		}

	}

	return newHead
}

// copyRandomList 随机链表的复制
// O(n) 空间复杂度
func copyRandomList2(head *Node) *Node {
	// 因为要深拷贝，需要两次遍历
	// 第一次遍历，遍历原节点和copyNode 的关系
	// 第二次遍历，根据原节点的Next和Random 建立链表关系

	hmap := make(map[*Node]*Node)

	// 第一次遍历：建立新节点，并放入hashMap
	// 要用i 复制一份，不然 head = head.Next， head 会到nil处
	for i := head; i != nil; {
		hmap[i] = &Node{i.Val, nil, nil}
		i = i.Next
	}

	// 第二次遍历:
	for i := head; i != nil; i = i.Next {
		copyed := hmap[i]
		copyed.Next = hmap[i.Next]
		copyed.Random = hmap[i.Random]
	}
	return hmap[head]
}
