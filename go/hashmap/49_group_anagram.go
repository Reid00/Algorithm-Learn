package hashmap

import "sort"

// groupAnagrams 将字母异位词分组，相同的放到一组
func groupAnagrams(strs []string) [][]string {
	// 排序，由于互为字母异位词，排序后 应该相同
	// 将排序后的字符串作为key 原始的字符串作为val 放到hashMap 中

	result := make([][]string, 0)

	hmap := make(map[string][]string)

	// 此处分组
	for _, str := range strs {
		b := []byte(str)

		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

		key := string(b)
		hmap[key] = append(hmap[key], str)
	}

	for _, v := range hmap {
		result = append(result, v)
	}
	return result
}

// groupAnagrams2 计数的方式来实现
// 用hashMap 记录每个字母的个数，然后用hashmap 做key -> hashmap 不能做key
// 因为只有26个小写字母，用[26]int 数组做key
func groupAnagrams2(strs []string) [][]string {
	result := make([][]string, 0)

	hmap := make(map[[26]int][]string)

	// 遍历分组
	for _, str := range strs {
		var key [26]int
		for _, ch := range str {
			k := ch - 'a' // 字符ch 转化为数组的下标
			key[k] += 1
		}

		hmap[key] = append(hmap[key], str)
	}

	for _, v := range hmap {
		result = append(result, v)
	}
	return result
}
