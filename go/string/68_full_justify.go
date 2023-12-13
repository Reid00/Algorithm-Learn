package string

import "strings"

// fullJustify 文本左右对齐 平均分布额外空格
// 填充空格的细节，我们分以下三种情况讨论:
// 1. 当前行是最后一行：单词左对齐，且单词之间应只有一个空格，在行末填充剩余空格；
// 2. 当前行不是最后一行，且只有一个单词：该单词左对齐，在行末填充空格；
// 3. 当前行不是最后一行，且不只一个单词：设当前行单词数为 numWords, 空格数为 numSpaces
// 我们需要将空格均匀分配在单词之间，则单词之间应至少有avgSpace=numSpaces/numWords -1
// 对于多出来的 extraSpaces=numSpaces % (numWords−1)个空格，
//  extraSpaces 应填在前 extraSpaces 个单词之间。
// 前 extraSpaces 个单词之间填充 avgSpaces+1 个空格，其余单词之间填充 avgSpaces 个空格。
func fullJustify(words []string, maxWidth int) []string {

	ans := make([]string, 0)

	i, n := 0, len(words)

	for {

		// words 的index i 作为新一行的开始
		start := i
		// 当前行字符的长度
		wordLen := 0
		// 每一行要尽量多的单词，同时该行的长度不能多于 maxWidth
		for ; i < n && wordLen+len(words[i])+i-start <= maxWidth; i++ {
			wordLen += len(words[i])
		}
		// 由于上面的循环在满足条件之后 i++ 了，所以i 的长度最大为n
		// 说明是最后一行的情况, 要左对齐，然后用空格填充
		if i == n {
			// 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
			s := strings.Join(words[start:i], " ")
			ans = append(ans, s+strings.Repeat(" ", maxWidth-len(s)))
			return ans
		}

		// 需要插入的空格数
		spaces := maxWidth - wordLen

		// 不是最后一行，判断是否是只有一个单词
		if i-start == 1 {
			// 只有一个单词
			s := words[start] + strings.Repeat(" ", spaces)
			ans = append(ans, s)
		} else {
			// 不是最后一行，不止一个单词的情况
			avg, extra := spaces/(i-start-1), spaces%(i-start-1)
			avgSpaces := strings.Repeat(" ", avg)

			// 前extra 单词之前用avgSpaces + 1 个单词填充
			s := strings.Join(words[start:start+extra+1], avgSpaces+" ")
			// extra 之后的词之前用avgSpaces 个单词填充
			s1 := strings.Join(words[start+extra+1:i], avgSpaces)
			ans = append(ans, s+avgSpaces+s1)
		}
	}

}
