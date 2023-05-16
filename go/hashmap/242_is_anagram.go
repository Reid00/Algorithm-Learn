package hashmap

// isAnagram 字母异位词
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	// 建立两个 hashMap 进行比较
	hmaps := make(map[rune]int)
	hmapt := make(map[rune]int)

	for _, v := range s {
		hmaps[v] += 1
	}

	for _, v := range t {
		hmapt[v] += 1
	}

	if len(hmaps) == len(hmapt) {
		for k := range hmaps {
			if hmaps[k] != hmapt[k] {
				return false
			}
		}
		return true
	}
	return false
}
