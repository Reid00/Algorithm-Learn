package string

// isSubsequence 判断子序列, s 是 t 的子序列
func isSubsequence(s string, t string) bool {
	// 双指针
	idx1, idx2 := 0, 0
	m, n := len(s), len(t)

	for idx1 < m && idx2 < n {
		if s[idx1] == t[idx2] {
			idx1 += 1
		}
		idx2 += 1
	}
	return idx1 == m
}
