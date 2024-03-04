package graph

import "slices"

// canFinish 课程表 判断是否可能完成所有课程的学习
// DAG 拓扑排序 DFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	// DFS
	// 当前节点是否被访问过, 搜索过的状态
	visited := make(map[int]bool)
	// 本次DFS 中是否被访问过， 搜索中的状态
	// 也可以用map 表示更方便，
	// 或者可以优化为都用vistied 表示，value 0 未访问，1访问中 + valid(bool) 2 已访问
	onStack := make([]int, 0)
	// 是否有环，如果有环无法形成拓扑排序
	hasCycle := false

	// 拓扑排序结果
	order := make([]int, 0)

	// build Graph
	graph := buildGraph(numCourses, prerequisites)

	var dfs func(int)

	dfs = func(u int) {
		// 同一次深度优先搜索时，当前顶点被访问过，说明存在环
		if slices.Contains(onStack, u) {
			hasCycle = true
		}
		if hasCycle || visited[u] {
			return
		}
		// 加入visited 中
		visited[u] = true
		// 标记本次深搜时，当前顶点被访问
		onStack = append(onStack, u)

		// dfs 其后继节点
		for _, v := range graph[u] {
			dfs(v)
		}

		// 遍历完之后回溯节点u
		order = append(order, u)
		// 移除正在搜索的状态
		onStack = slices.DeleteFunc(onStack, func(i int) bool { return i == u })
	}

	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			dfs(i)
		}
	}
	return !hasCycle
}

// buildGraph 构建图结构
func buildGraph(numCourse int, prereq [][]int) [][]int {

	graph := make([][]int, numCourse)

	// 初始化邻接表
	for i := 0; i < numCourse; i++ {
		graph[i] = make([]int, 0)
	}

	// 遍历每个点的先后关系，构建Graph, 节点 -> 节点的邻居节点
	for _, rel := range prereq {
		graph[rel[1]] = append(graph[rel[1]], rel[0])
	}

	return graph
}

// canFinish
// Kahn 算法 根据图结构，从入度 inDegree 为0 的开始遍历
// 遍历完0 的节点，如果indegree==0, 将其加入拓扑排序order 中，然后遍历 其出度，并将其出度入度-1
// 重复上述步骤，如果len(order) == len(inDegree) 说明无环 可以完成拓扑排序
func canFinish2(numCourses int, prerequisites [][]int) bool {

	graph := buildGraph(numCourses, prerequisites)

	// 记录每个节点的入度
	inDegree := make([]int, numCourses)

	// 拓扑排序结果
	order := make([]int, 0)

	// 计算入度
	for i := 0; i < numCourses; i++ {
		// 节点j 的入度
		for j := 0; j < len(graph[i]); j++ {
			inDegree[graph[i][j]] += 1
		}
	}

	// zeroDegree 记录入度为0 的节点
	zeroDegree := make([]int, 0)
	for i, v := range inDegree {
		if v == 0 {
			zeroDegree = append(zeroDegree, i)
		}
	}

	for len(zeroDegree) > 0 {

		for _, v := range zeroDegree {
			// 出队列
			zeroDegree = zeroDegree[1:]
			order = append(order, v)
			// 节点v 是入度为0 的节点, 将其邻居节点入度-1
			for _, v := range graph[v] {
				inDegree[v]--
				if inDegree[v] == 0 {
					zeroDegree = append(zeroDegree, v)
				}
			}

		}
	}

	return len(order) == numCourses
}
