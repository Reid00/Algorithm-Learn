package graph

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph 克隆图  深拷贝（克隆）一个图
// BFS
func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}

	// 哈希表中的 key 是原始图中的节点，value 是克隆图中的对应节点。
	visited := make(map[*Node]*Node)

	var dfs func(*Node) *Node

	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		// 如果已经访问过了，直接返回该节点的克隆
		if _, ok := visited[n]; ok {
			return visited[n]
		}
		// deepcopy node
		clonedNode := &Node{Val: n.Val, Neighbors: make([]*Node, 0)}

		// 记录已经访问过的节点，避免出现循环
		visited[n] = clonedNode

		// 填充其邻居节点
		for _, neighbor := range n.Neighbors {
			clonedNode.Neighbors = append(clonedNode.Neighbors, dfs(neighbor))
		}

		return clonedNode

	}
	return dfs(node)
}

// cloneGraph 克隆图  深拷贝（克隆）一个图
// BFS
func cloneGraph2(node *Node) *Node {

	if node == nil {
		return nil
	}
	// dfs 和 BFS 本质差不多，此处依旧需要使用hmap 记录每个节点是否被访问克隆

	visited := make(map[*Node]*Node)
	visited[node] = &Node{
		Val:       node.Val,
		Neighbors: nil,
	}
	q := make([]*Node, 0)
	q = append(q, node)

	for len(q) > 0 {

		for _, n := range q {
			// 出队列
			q = q[1:]

			// 确认邻居节点是否被clone
			for _, neighbor := range n.Neighbors {
				if _, ok := visited[neighbor]; !ok {
					visited[neighbor] = &Node{
						Val: neighbor.Val,
					}

					q = append(q, neighbor)
				}
				// 更新邻居列表
				visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbor])
			}

		}

	}
	return visited[node]
}
