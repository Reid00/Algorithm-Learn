package main

// 单源节点的最短路径

import (
	"fmt"
	"math"
)

func main() {

	prerequisites := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 2
	res := networkDelayTime(prerequisites, n, k)
	fmt.Println(res)

}

func networkDelayTime(times [][]int, n int, k int) int {
	// Dijkstra 算法是核心
	// 思想，在图中，找到为访问的距离源点最近的点，并标记为访问 == 找出未访问过的到原点距离最小的顶点并标记为访问
	// 步骤：1. 找到距离源点最近的点
	// 2. 标记为访问
	// 3. 重复1，2

	// 题解思路
	// 1. 用邻接矩阵 构建Graph
	// 2. 初始化distance 切片
	// 3. 初始化isUsed 切片
	// 4. 用Dijkstra 算法 遍历所有的点，更新distance
	// 5. 遍历distance 切片，如果存在inf 值，则说明有源点到达不了的点，返回-1 否则返回最大值

	const inf = math.MaxInt64 / 2
	// 1. init Graph
	graph := buildGraph(times, n)

	// 2. init distance
	distance := make([]int, n)
	// 源点k 距离设为0， 其余为inf
	for i := range distance {
		distance[i] = inf
	}
	distance[k-1] = 0

	// 3. init isUsed
	isUsed := make([]bool, n)

	// Dijkstra 算法
	for i := 0; i < n; i++ {
		// 根据距离和是否遍历过，找出距离源点最近的节点，第一次从源点k-1 开始
		pick := -1

		for i, used := range isUsed {
			if !used && (pick == -1 || distance[i] < distance[pick]) { //遍历isUsed 的同时，遍历distance 找出距离最小的节点
				pick = i
			}
		}
		isUsed[pick] = true

		// 4. 更新距离
		for node, v := range graph[pick] {
			// distance[node] 表示node节点到源点的距离
			distance[node] = min(distance[node], distance[pick]+v) //distance[pick] + v 是一条新的找源点的路径值
		}

	}
	ans := 0
	// 5. 遍历distance 找出最大的dis
	for _, v := range distance {
		if v == inf {
			return -1
		}
		ans = max(ans, v)
	}
	return ans
}

func buildGraph(times [][]int, n int) [][]int {
	const inf = math.MaxInt64 / 2
	// n * n 的矩阵
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
		// 初始化默认距离为inf
		for j := range graph[i] {
			graph[i][j] = inf
		}
	}

	// 根据times 修改对应的距离
	for i := range times {
		// 因为times 中节点id 从1 开始，而邻接矩阵中，节点id 从0 开始，所以row, col 要在times 中减一
		row, col := times[i][0]-1, times[i][1]-1
		graph[row][col] = times[i][2]
	}

	return graph
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
