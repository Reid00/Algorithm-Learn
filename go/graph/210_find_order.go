package graph

import "slices"

// findOrder 课程表 II 返回你为了学完所有课程所安排的学习顺序
// DAG DFS
func findOrder(numCourses int, prerequisites [][]int) []int {

	graph := buildGraph(numCourses, prerequisites)

	// 每个节点的状态，是否被访问过
	visited := make(map[int]bool)

	// 每个节点本次DFS 中是否是正在搜索中的状态
	onSearch := make(map[int]bool)

	// order 拓扑排序的结果
	order := make([]int, 0)

	// 是否有环
	hasCycle := false

	var dfs func(int)

	dfs = func(u int) {

		// 本次DFS 中已经被标记为搜索
		if onSearch[u] {
			hasCycle = true
		}

		// 有环、已经被搜索
		if hasCycle || visited[u] {
			return
		}
		// 标记为本次DFS 搜索中
		onSearch[u] = true
		visited[u] = true

		// 遍历u 的邻居节点
		for _, v := range graph[u] {
			dfs(v)
		}

		// 回溯，将u 入栈到order 中 (注意此处压栈时先入栈的是u 的叶子节点)
		order = append(order, u)
		// 将u 的状态由onSearch 中移除
		delete(onSearch, u)
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	if hasCycle {
		return nil
	}
	// order 的顺序是逆序的，先是叶子节点才是root
	slices.Reverse(order)
	return order
}
