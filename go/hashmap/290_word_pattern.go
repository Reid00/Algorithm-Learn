package hashmap

import "strings"

// wordPattern 单词规律
func wordPattern(pattern string, s string) bool {
	// HashMap
	// 建立两个hashmap pattern 中的字符 -> s 中的string 的映射 和 反过来的映射

	hmap1 := make(map[byte]string)
	hmap2 := make(map[string]byte)

	vals := strings.Split(s, " ")

	if len(vals) != len(pattern) {
		return false
	}

	for i, v := range pattern {

		if val, ok := hmap1[byte(v)]; ok {
			if val != vals[i] {
				return false
			}
		}

		if val, ok := hmap2[vals[i]]; ok {
			if val != byte(v) {
				return false
			}
		}

		hmap1[byte(v)] = vals[i]
		hmap2[vals[i]] = byte(v)

	}
	return true
}
