package trie

// 用字典树实现
// word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母
// 因为.可以代表任意子母，所以遍历到. 的时候，要查看其所有孩子里面是否有满足的
type WordDictionary struct {
	children [26]*WordDictionary
	end      bool
}

func Constructor2() WordDictionary {
	return WordDictionary{}
}

// AddWord 新家一个单词
func (w *WordDictionary) AddWord(word string) {

	node := w
	for _, v := range word {
		idx := v - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &WordDictionary{}
		}
		node = node.children[idx]
	}

	node.end = true
}

// Search 搜索一个单词是否存在
func (w *WordDictionary) Search(word string) bool {

	// DFS
	var dfs func(int, *WordDictionary) bool

	dfs = func(i int, wd *WordDictionary) bool {
		if i == len(word) {
			return wd.end
		}

		ch := word[i]

		if ch != '.' {
			idx := ch - 'a'
			child := wd.children[idx]
			// 继续往下遍历孩子节点
			if child != nil && dfs(i+1, child) {
				return true
			}
		} else {

			// . 可以代表 所有字符，遍历wd 的所有孩子
			for _, v := range wd.children {
				if v != nil && dfs(i+1, v) {
					return true
				}
			}
		}

		return false
	}

	return dfs(0, w)
}
