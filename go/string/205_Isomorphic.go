package string

// isIsomorphic 判断两个字符串是否为同构
func isIsomorphic(s string, t string) bool {

	s2t := make(map[byte]byte) // s 转 t 记录情况
	t2s := make(map[byte]byte) // t 转 s 记录情况

	for i := range s {
		x, y := s[i], t[i]

		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}

		s2t[x] = y // NOTE x 转 y
		t2s[y] = x
	}

	return true
}
