package string

import "strings"

func lengthOfLastWord(s string) int {

	s = strings.TrimSpace(s)

	result := strings.Split(s, " ")
	ret := result[len(result)-1]

	return len(ret)
}

// lengthOfLastWord2 从后往前遍历
func lengthOfLastWord2(s string) int {

	var ans int
	index := len(s) - 1
	// 去除末尾的空格
	for s[index] == ' ' {
		index--
	}

	// 上面找到了最后一个单词的最后一个字符，从最后的字符开始遍历
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}

	return ans
}
