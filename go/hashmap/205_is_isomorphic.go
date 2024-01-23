package hashmap

//  isIsomorphic 同构字符串
func isIsomorphic(s string, t string) bool {
	// 维护两个hashMap 从s->t  和  t-> s

	s2t, t2s := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sCh, tCh := s[i], t[i]

        // s2t[sCh] >0 是为了保证s2t 先存在
        // <=0  不存在，先建立映射
		if s2t[sCh] >0 && s2t[sCh] != tCh || t2s[tCh] >0 && t2s[tCh] != sCh {
			return false
		}

		s2t[sCh] = tCh
		t2s[tCh] = sCh
	}
	return true
}