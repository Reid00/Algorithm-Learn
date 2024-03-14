package divideandconquer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {

	return helper(nums, 0, len(nums)-1)
}

// 递归，总是选择中间元素 靠左的座位跟节点
func helper(nums []int, left, right int) *TreeNode {

	if left > right {
		return nil
	}

	mid := left + (right-left)>>1
	root := &TreeNode{Val: nums[mid]}

	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root
}
