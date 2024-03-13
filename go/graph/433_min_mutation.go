package graph

import (
	"slices"
)

// minMutation 最小基因变化
// BFS
func minMutation(startGene string, endGene string, bank []string) int {

	if !slices.Contains(bank, endGene) {
		return -1
	}

	if startGene == endGene {
		return 0
	}

	var mutationMap = map[byte][3]string{
		'A': [...]string{"T", "G", "C"},
		'C': [...]string{"T", "G", "A"},
		'T': [...]string{"A", "G", "C"},
		'G': [...]string{"T", "A", "C"},
	}

	// 记录bank 中的元素是否 转化过
	visited := make([]bool, len(bank))

	queue := make([]string, 0)
	queue = append(queue, startGene)

	// 转化的步骤
	cnt := 0

	for len(queue) > 0 {

		for _, v := range queue {
			queue = queue[1:]

			if v == endGene {
				return cnt
			}

			for j := 0; j < len(v); j++ {
				// 对字符进行转化
				for _, s := range mutationMap[v[j]] {
					// 将第j 个字符替换为map 中的一个
					newGene := v[:j] + s + v[j+1:]
					// 突变到一个没有试过的基因，记录下来
					if idx := slices.Index(bank, newGene); idx != -1 && !visited[idx] {
						queue = append(queue, bank[idx])
						visited[idx] = true
					}

				}
			}
		}
		cnt++
	}

	return -1
}

// 双向BFS
func minMutation2(startGene string, endGene string, bank []string) int {

	// 为了防止最后一层维度爆炸，双向BFS 来回交替遍历入队队列
	// 这里的做法就是将起点和终点分别加入两个集合当中，然后从集合元素少的一端开始搜索，.
	// 这样能够减少搜索量，直到两个集合产生了交集，那么就结束搜索。

	if startGene == endGene {
		return 0
	}

	if slices.Index(bank, endGene) == -1 {
		return -1
	}

	mutationMap := map[byte][3]string{
		'A': [...]string{"T", "G", "A"},
		'C': [...]string{"T", "G", "A"},
		'T': [...]string{"A", "G", "C"},
		'G': [...]string{"T", "A", "C"},
	}

	// 记录bank 中的基因是否已经被访问过
	visited := make([]bool, len(bank))

	startQ := make([]string, 0)
	startQ = append(startQ, startGene)
	endQ := make([]string, 0)
	endQ = append(endQ, endGene)

	var cnt int

	for len(startQ) > 0 {

		cnt++

		for _, v := range startQ {
			startQ = startQ[1:]

			// 对 v 中的每个字符尝试突变
			for i := 0; i < len(v); i++ {

				for _, s := range mutationMap[v[i]] {

					newGene := v[:i] + s + v[i+1:]

					// 两个集合有交集代表可以找到
					if idx := slices.Index(endQ, newGene); idx != -1 {
						return cnt
					}

					if idx := slices.Index(bank, newGene); idx != -1 && !visited[idx] {
						startQ = append(startQ, newGene)
						visited[idx] = true
					}
				}
			}

		}
		// 从搜索空间小的Q 进行搜索
		if len(startQ) > len(endQ) {
			startQ, endQ = endQ, startQ
		}

	}

	return -1
}
