/*
@Time        :2022/03/15 17:00:41
@Author      :Reid
@Version     :1.0
@Desc        :前缀树的实现
*/

package tree

// Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
// 请你实现 Trie 类：
// Trie() 初始化前缀树对象。
// void insert(String word) 向前缀树中插入字符串 word 。
// boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
// boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。

type Trie struct {
	child map[byte]*Trie
	isEnd bool
}

func NewTrie() *Trie {
	return &Trie{}
}

// 字典树中插入字符串
func (t *Trie) Insert(word string) {

	for _, v := range word {
		b := byte(v) // word 都是小写英文 肯定可以转化成byte 类型

		if t.child == nil { // 不存在child存储的hashMap时
			t.child = make(map[byte]*Trie)
		}
		if _, ok := t.child[b]; !ok { // 不存在b字符
			t.child[b] = NewTrie()
		}
		t = t.child[b]
	}
	t.isEnd = true
}

// 搜索前缀
func (t *Trie) searchPrefix(word string) *Trie {

	for _, v := range word {
		b := byte(v)
		if t.child[b] == nil { // 不存在
			return nil
		}
		t = t.child[b]
	}

	return t
}

// 搜索单次是否存在
func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(word string) bool {
	return t.searchPrefix(word) != nil
}
