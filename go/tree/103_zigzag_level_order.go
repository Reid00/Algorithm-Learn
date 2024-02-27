package tree

// zigzagLevelOrder 二叉树的锯齿形层序遍历
// 引入flag，偶数从左到右，奇数数从右到左
func zigzagLevelOrder(root *TreeNode) [][]int {
    // 层次遍历，引入flag 在奇数层 对层级结果反转

    if root == nil {
        return nil
    }

    queue := make([]*TreeNode, 0)    
    queue = append(queue, root)

    ans := make([][]int, 0)

    var flag int

    for len(queue) > 0 {

        level := make([]int, 0)
        for _, v := range queue {
            queue = queue[1:]
            level = append(level, v.Val)

            if v.Left != nil {
                queue =append(queue, v.Left)
            }
            if v.Right != nil {
                queue = append(queue, v.Right)
            }
        }
        // 奇数反转
        if flag & 1 == 1{
            reverse(level)
        }
        flag ++
        ans = append(ans, level)
    }
    return ans 
}

// zigzagLevelOrder 二叉树的锯齿形层序遍历
// 引入flag，偶数从左到右，奇数数从右到左
func zigzagLevelOrder2(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	ans := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	// false default 不用reverse
	var flag bool

	for len(queue) > 0 {
		level := make([]int, 0)

		length := len(queue)

		for i := 0; i < length; i++ {
			head := queue[0]
			queue = queue[1:]

			level = append(level, head.Val)
			if head.Left != nil {
				queue = append(queue, head.Left)
			}

			if head.Right != nil {
				queue = append(queue, head.Right)
			}
		}
		if flag {
			reverse(level)
		}

		flag = !flag
		ans = append(ans, level)
	}

	return ans
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
