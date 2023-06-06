package unionfind

// equationsPossible 用并查集构建连通分量来判断equations 中的== 和 != 是否存在矛盾
func equationsPossible(equations []string) bool {
	// equations 是小写字母，所以连通分量的数组长度是26
	uf := NewUnionFind(26)

	// 输入：["a==b","b!=a"]
	// 输出：false
	// == 意味着连通到一起
	for _, str := range equations {
		//  如果== 则连通到一起
		if str[1] == '=' {
			idx1 := str[0] - 'a'
			idx2 := str[3] - 'a'

			uf.Union(int(idx1), int(idx2))
		}
	}

	// 处理 != 的情况
	for _, str := range equations {
		if str[1] == '!' {
			idx1 := str[0] - 'a'
			idx2 := str[3] - 'a'

			if uf.Find(int(idx1)) == uf.Find(int(idx2)) {
				return false
			}
		}
	}
	return true
}
