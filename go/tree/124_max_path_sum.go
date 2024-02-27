package tree

import "math"

// maxPathSum 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {

	// 最大路径和
	maxSum := math.MinInt32
	// 计算当前节点的最大贡献
	// 节点>0 val 是贡献值，<0 贡献为0
	var maxGain func(*TreeNode) int

	maxGain = func(tn *TreeNode) int {

		if tn == nil {
			return 0
		}

		// 左右节点的贡献值
		leftGain := maxGain(tn.Left)
		rightGain := maxGain(tn.Right)

		// innerMaxSum 只有当tn 是root 的时候才会同时去左右子树，否则只能去一个方向
		innerMaxSum := leftGain + tn.Val + rightGain
		// 更新最大路径和
		maxSum = max(maxSum, innerMaxSum)

		// tn 不是root 的时候对外提供的受益
		outMaxSum := tn.Val + max(leftGain, rightGain) // maxGain 返回是>=0的，所以不需要考虑负数

		return max(outMaxSum, 0)
	}

	maxGain(root)
	return maxSum
}
