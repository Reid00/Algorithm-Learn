package string

import "strings"

// convert N 字形变换
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	// 1. 有几行创建几个元素
	// 2. 将s 里面的字符遍历(即是N字形状，先竖着在上去)依次按照 写入res[i]
	// 然后到达 numRows-1 之后 倒序写入res中 即numRows-2
	res := make([]string, numRows)

	i, flag := 0, -1

	for _, ch := range s {

		res[i] += string(ch)
		// 开始的时候N字形 向下写入
		// 到底部的时候倒序写入 res
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		// flag 一直是1 或者-1
		i += flag
	}

	return strings.Join(res, "")

}
