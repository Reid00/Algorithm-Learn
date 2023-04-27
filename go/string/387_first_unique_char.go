package string

// firstUniqChar 字符串中，第一个不重复字符
func firstUniqChar(s string) int {

	// 思路: hmap + array
	// hmap 确保是否存在重复， array 记录是否顺序，保证第一

	// 优化: 遍历两次, 不用array 存储
	// 计数不能用byte, 会越界
	hmap := make(map[rune]int)

	for _, ch := range s {
		hmap[ch] += 1
	}

	for i, ch := range s {
		if hmap[ch] == 1 {
			return i
		}
	}
	return -1
}
