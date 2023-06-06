package unionfind

// findCircleNum 省份数量
func findCircleNum(isConnected [][]int) int {

	n := len(isConnected)

	// parent[x] 代表节点x 的父节点是parent[x]
	parent := make([]int, n)

	// 节点i 的父节点 初始指向自己
	for i := range parent {
		parent[i] = i
	}

	// find 找到x 的父节点
	var find func(x int) int
	find = func(x int) int {
		// 路径压缩技巧，把找到x 的父节点，然后把x 到父节点之间的所有节点，放到父节点下面
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(from, to int) {
		// union 将from 和 to 连接到一起，即将其根节点连接到一起
		rootFrom := find(from)
		rootTo := find(to)
		if rootFrom == rootTo {
			return
		}
		parent[from] = rootTo
		// 或者写 一样的效果
		// parent[to] = rootFrom
	}

	// 遍历邻接表
	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 {
				union(i, j)
			}
		}
	}

	var ans int
	for i, p := range parent {
		// i == p 意味着i 的父节点等于自己，即是根节点的情况
		if i == p {
			ans++
		}
	}

	return ans
}
