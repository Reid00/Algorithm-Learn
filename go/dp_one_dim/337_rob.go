package dponedim

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 打家劫舍 III
// 房屋邻接是树形，不是数组，所以需要dfs
func rob3(root *TreeNode) int {
	// 遍历每个节点的时候 都有两种选择，偷当前节点和不偷当前节点，所以用个长度为2的数组dp即可
	// 下标为0记录不偷该节点所得到的的最大金钱，下标为1记录偷该节点所得到的的最大金钱。
	// 偷当前节点的时候不能偷左右孩子
	// 遍历顺序, 用后序遍历
	// 通过递归左节点，得到左节点偷与不偷的金钱。	left = robTree(root.Left)
	// 	通过递归右节点，得到右节点偷与不偷的金钱。  right = robTree(root.Right)
	// 如果是偷当前节点，那么左右孩子就不能偷，val1 = cur->val + left[0] + right[0];
	// 如果不偷当前节点，那么左右孩子就可以偷，至于到底偷不偷一定是选一个最大的，
	// 所以：val2 = max(left[0], left[1]) + max(right[0], right[1]);

	if root == nil {
		return 0
	}

	var robTree func(*TreeNode) []int
	robTree = func(node *TreeNode) []int {
		// 长度为2 的dp
		res := make([]int, 2)
		if node == nil {
			return res
		}

		// 递归左节点，得到做节点偷与不偷得到的金钱
		left := robTree(node.Left)
		// 递归右节点
		right := robTree(node.Right)

		// 偷cur 节点，则不能偷左右节点
		val1 := node.Val + left[0] + right[0]
		res[1] = val1
		// 不偷cur 节点，在左右子孩子中选择偷或者不偷
		val2 := max(left[0], left[1]) + max(right[0], right[1])
		res[0] = val2
		// 注意0 是不偷，1是偷当前节点，别弄反了
		return res
	}

	val := robTree(root)

	return max(val[0], val[1])
}
