package tree

// 二叉搜索树迭代器
// 实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
// BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
// boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
// int next()将指针向右移动，然后返回指针处的数字。
type BSTIterator struct {
	stack []*TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		stack: make([]*TreeNode, 0),
		cur:   root,
	}
}

// 通过迭代的方式对二叉树做中序遍历。此时，我们无需预先计算出中序遍历的全部结果，只需要实时维护当前栈的情况即可。
func (b *BSTIterator) Next() int {

	// 先左子树
	for node := b.cur; node != nil; node = node.Left {
		b.stack = append(b.stack, node)
	}
	// 出栈
	b.cur, b.stack = b.stack[len(b.stack)-1], b.stack[:len(b.stack)-1]

	val := b.cur.Val
	// 右子树
	b.cur = b.cur.Right
	return val
}

func (b *BSTIterator) HasNext() bool {
	return b.cur != nil || len(b.stack) > 0
}
