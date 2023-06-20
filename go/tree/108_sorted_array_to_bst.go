package tree

// sortedArrayToBST 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {

	// 中序遍历，以中间节点作为根节点

	return builder(nums, 0, len(nums)-1)
}

func builder(nums []int, start, end int) *TreeNode {

	// 中序遍历，以中间节点作为根节点
	if start > end {
		return nil
	}

	// 奇数根节点为中间点，偶数 为中间节点偏左的
	// not use(start + end) / 2 防止int 越界
	mid := start + (end-start)>>1

	root := &TreeNode{
		Val: nums[mid],
	}

	// BUG this root.Left = builder(nums, 0, mid-1)
	root.Left = builder(nums, start, mid-1)
	root.Right = builder(nums, mid+1, end)
	return root
}
