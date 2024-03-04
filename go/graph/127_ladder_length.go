package graph

import "slices"

// ladderLength 单词接龙
func ladderLength(beginWord string, endWord string, wordList []string) int {

	// 双向BFS， 减少搜索空间
	if beginWord == endWord {
		return 0
	}

	if idx := slices.Index(wordList, endWord); idx == -1 {
		return 0
	}

	// 每个节点是否被转化过
	visited := make([]bool, len(wordList))

	// 双向BFS， 两个队列
	startQ := []string{beginWord}
	endQ := []string{endWord}

	var cnt int

	for len(startQ) > 0 {

		cnt++

		for _, v := range startQ {

			startQ = startQ[1:]

			for i := 0; i < len(v); i++ {

				for s := 'a'; s <= 'z'; s++ {

					newWord := v[:i] + string(s) + v[i+1:]

					// 两个队列有交集则可以找到
					if idx := slices.Index(endQ, newWord); idx != -1 {
						return cnt + 1
					}

					if idx := slices.Index(wordList, newWord); idx != -1 && !visited[idx] {
						startQ = append(startQ, newWord)
						visited[idx] = true
					}

				}

			}
		}

		// 减少搜索空间，只搜索较小的空间
		if len(startQ) > len(endQ) {
			startQ, endQ = endQ, startQ
		}

	}

	return 0
}

// ladderLength 单词接龙
func ladderLength2(beginWord string, endWord string, wordList []string) int {
	// 单向BFS， 最短路径用BFS
	// 某个test cases 超时

	if beginWord == endWord {
		return 0
	}

	if slices.Index(wordList, endWord) == -1 {
		return 0
	}

	// 节点状态是否访问过
	visited := make([]bool, len(wordList))

	queue := make([]string, 0)
	queue = append(queue, beginWord)

	var cnt int

	for len(queue) > 0 {

		for _, v := range queue {
			queue = queue[1:]

			if v == endWord {
				// 返回不是变化次数，是序列长度
				return cnt + 1
			}

			// v 中的每个字符都有可能发生变化
			for i := 0; i < len(v); i++ {

				for s := 'a'; s <= 'z'; s++ {

					newWord := v[:i] + string(s) + v[i+1:]

					if idx := slices.Index(wordList, newWord); idx != -1 && !visited[idx] {
						queue = append(queue, newWord)
						visited[idx] = true
					}
				}
			}
		}

		cnt++

	}
	return 0
}
