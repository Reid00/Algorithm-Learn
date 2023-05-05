package string

import "reflect"

// findAnagrams 找到字符串中所有字母异位词
// 异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
// 暴力破解 未通过
func findAnagrams(s string, p string) []int {

	n := len(p)
	if n == 0 {
		return []int{}
	}
	result := make([]int, 0)
	// hashMap 保存字符的个数，相同则是字母异味词
	hashP := make(map[rune]int)

	for _, ch := range p {
		hashP[ch] += 1
	}

	for i, ch := range s {
		if hashP[ch] > 0 {
			hashSub := make(map[rune]int)
			if i+n <= len(s) {
				for _, v := range s[i : i+n] {
					hashSub[v] += 1
				}
			}
			if reflect.DeepEqual(hashSub, hashP) {
				result = append(result, i)
			}
		}
	}
	return result
}

// findAnagrams2 滑动窗口, 记录窗口里面的字符有哪些
func findAnagrams2(s string, p string) []int {
	ans := make([]int, 0)
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}

	var sCnt, pCnt [26]int
	for i, ch := range p {
		sCnt[s[i]-'a']++
		pCnt[ch-'a']++
	}

	if sCnt == pCnt {
		// 开头就相同的异位字母，index 0
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		// 遍历到ch 意味着ch 要被去掉
		sCnt[ch-'a']--
		// i+1 开始，长度为pLen 观察是否相同
		sCnt[s[i+pLen]-'a']++
		if sCnt == pCnt {
			ans = append(ans, i+1)
		}
	}
	return ans
}

// findAnagrams3 滑动窗口, 记录窗口里面的字符差异
func findAnagrams3(s string, p string) []int {
	ans := make([]int, 0)
	sLen, pLen := len(s), len(p)
	if pLen > sLen {
		return ans
	}

	// diff window 记录s 和 p的字符差异
	// diffCnt 记录差异字符的个数
	diff := [26]int{}
	diffCnt := 0
	for i, ch := range p {
		// 记录原始串中字符的分布
		diff[s[i]-'a']++
		// 遍历p 没遍历一个字符就在对应位置删除一个
		diff[ch-'a']--
	}

	for _, v := range diff {
		if v != 0 {
			diffCnt++
		}
	}

	if diffCnt == 0 {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		// 遍历到ch， 意味着要去掉ch
		if diff[ch-'a'] == 1 {
			// 原先ch 相差1， 现在窗口中离开，由不同变的相同
			diffCnt--
		} else if diff[ch-'a'] == 0 {
			// 由相同变为不同
			diffCnt++
		}
		// diff window中，ch个数要减一
		diff[ch-'a']--

		// 新加入一个字符
		// i+1 开始，长度为pLen 观察是否相同
		if diff[s[i+pLen]-'a'] == -1 {
			diffCnt--
		} else if diff[s[i+pLen]-'a'] == 0 {
			diffCnt++
		}

		diff[s[i+pLen]-'a']++
		if diffCnt == 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
