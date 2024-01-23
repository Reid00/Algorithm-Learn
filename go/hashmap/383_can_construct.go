package hashmap

// canConstruct 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	// HashMap
	if len(magazine) == 0 || len(magazine) < len(ransomNote) {
		return false
	}

	hmap := make(map[rune]int)

	for _, v := range magazine {
		hmap[v]++
	}

	for _, v := range ransomNote {
		// if cnt, ok := hmap[v]; ok {
		// 	if cnt <= 0 {
		// 		return false
		// 	} else {
		// 		hmap[v]--
		// 	}
		// } else {
		// 	return false
		// }

		hmap[v]--
		if hmap[v] < 0 {
			return false
		}
	}
	return true
}
