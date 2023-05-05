package string

import "strings"

// reverseWords 反转单词, call string API
// 输入： s = "God Ding"
// 输出："doG gniD"
func reverseWords4(s string) string {
	words := strings.Fields(s)

	result := make([]string, 0)

	for _, word := range words {
		reversedWord := reverseStr2(word)
		result = append(result, reversedWord)
	}

	return strings.Join(result, " ")
}

func reverseStr2(s string) string {
	b := []byte(s)

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func reverseWords5(s string) string {
	// 转成byte 之后，双指针
	b := []byte(s)

	var slow, fast int

	for fast < len(b) {
		// 考虑最后一个单词时， b[fast] != ' '
		if fast == len(b)-1 {
			reverseWithIndex(b, slow, fast)
			break
		}

		if b[fast] != ' ' {
			fast++
			continue
		}

		// 到此处说明,fast 所在字符为空格
		reverseWithIndex(b, slow, fast-1)
		slow = fast + 1
		fast = fast + 1
	}

	return string(b)
}
