package trie

// 实现字典树 Trie
type Trie struct {
	children [26]*Trie
	end      bool
	// 额外信息，用于加速 leetcode 212 题
	count int    // 记录本节点被覆盖多少次，用于剪枝
	word  string //保存整个路径，便于直接扔到本题的结果数组中。否则就得在dfs过程中额外维护路径信息
}

func Constructor() Trie {
	return Trie{}
}

// Insert 插入一个单词
func (t *Trie) Insert(word string) {

	for _, ch := range word {
		idx := ch - 'a'

		if t.children[idx] == nil {
			t.children[idx] = &Trie{}
		}

		t = t.children[idx]
		// 该节点被使用次数
		t.count++
	}
	t.end = true
	t.word = word
}

// Del 剪枝 for leetcode 212
func (t *Trie) Del(word string) {
	for _, ch := range word {
		idx := ch - 'a'
		t = t.children[idx]
		t.count--
	}
	t.end = false
}

// Search 搜索一个单词是否存在
func (t *Trie) Search(word string) bool {
	trie := t.SearchPrefix(word)
	return trie != nil && trie.end
}

// StartsWith 是否以某个单词为前缀
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}

func (t *Trie) SearchPrefix(prefix string) *Trie {

	for _, ch := range prefix {
		idx := ch - 'a'
		if t.children[idx] == nil {
			return nil
		}

		t = t.children[idx]
	}
	return t
}
