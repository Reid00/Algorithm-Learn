package main

//层次遍历
func levelTraversal(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		length := len(queue)

		list := make([]int, 0)

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, list)

	}

	return result
}

// 解开密码锁问题，Leetcode 752
func openLock(deadends []string, target string) int {
	start := "0000"

	if target == start {
		return 0
	}

	dead := make(map[string]bool, 0)
	// dead := map[string]bool{}
	// 哈希保存死锁的字符串
	for _, v := range deadends {
		dead[v] = true
	}
	// 如果死锁发生在开始的时候，返回-1
	if dead[start] {
		return -1
	}

	// 从start 开始， 生成有可能产生的候选结果构建成虚拟的多叉树
	genCandidate := func(start string) (ret []string) {
		s := []byte(start)

		// 产生候选字符串
		for i, v := range s {
			s[i] = v + 1

			if s[i] > '9' { // 注意: byte 比较
				s[i] = '0'
			}
			ret = append(ret, string(s))

			s[i] = v - 1
			if s[i] < '0' {
				s[i] = '9'
			}
			ret = append(ret, string(s))
			// 重要: 每次修改完成之后都要修改回去
			s[i] = v
		}
		return
	}

	// BFS 找到最短路径到target
	type pair struct {
		val  string
		step int
	}

	queue := make([]*pair, 0)
	queue = append(queue, &pair{val: start, step: 0})

	// 定义元素是否被访问过
	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0 {
		length := len(queue)

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			candidates := genCandidate(node.val)
			for _, v := range candidates {
				if !visited[v] && !dead[v] {
					if v == target {
						return node.step + 1
					}
					visited[v] = true
					nextLevel := &pair{
						val:  v,
						step: node.step + 1,
					}
					queue = append(queue, nextLevel)
				}
			}

		}
	}

	return -1
}
