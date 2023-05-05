package string

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)

	splited := strings.Split(s, " ")
	var result string

	for i := len(splited) - 1; i >= 0; i-- {
		if strings.TrimSpace(splited[i]) == "" {
			continue
		}
		if i == 0 {
			result += splited[i]
		} else {
			result += splited[i] + " "
		}
	}
	return result
}

func reverseWords2(s string) string {

	words := strings.Fields(s)

	resultSli := make([]string, 0, len(words))

	for i := range words {
		resultSli = append(resultSli, words[len(words)-1-i])
	}

	return strings.Join(resultSli, " ")
}

// reverseWords3 自己实现，不调用API， 空间复杂度为O(1)
func reverseWords3(s string) string {

	// covert to bytes
	b := []byte(s)
	// 使用快慢指针，去除指定的字符，本处是空格
	// 1. 先去除头部的空格
	var slow, fast int
	for len(b) > 0 && b[fast] == ' ' && fast < len(b) {
		fast++
	}

	// 2. 去除中间冗余(连续多个空格的情况，保留一个空格)的空格
	for ; fast < len(b); fast++ {
		if fast-1 > 0 && b[fast] == b[fast-1] && b[fast] == ' ' {
			continue
		}

		// 快慢指针，将非空格移到slow index 处
		// 这样，slow 左侧都是满足条件的字符
		b[slow] = b[fast]
		slow++
	}

	// 3. 去除尾部的空格, 在上面slow + 1， 所以最后一个字符index 是slow -1
	if slow-1 > 0 && b[slow-1] == ' ' {
		b = b[:slow-1]
	} else {
		b = b[:slow]
	}

	// 反转整个字符串
	reverseWithIndex(b, 0, len(b)-1)

	// 反转单词, 以空格为分隔符区分单词
	i := 0
	for i < len(b) {
		j := i

		for ; j < len(b) && b[j] != ' '; j++ {
			// 直到碰到空格才推出循环
		}
		// 反转单词
		reverseWithIndex(b, i, j-1)
		i = j + 1
	}

	return string(b)
}

// reverseWithIndex 反转b[start: end+1], 包含end
func reverseWithIndex(b []byte, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
