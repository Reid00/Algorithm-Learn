package string

import "strings"

// toLowerCase API
func toLowerCase(s string) string {
	return strings.ToLower(s)
}

// toLowerCase2 自己实现
// 大写变小写、小写变大写：字符 ^= 32 （大写 ^= 32 相当于 +32，小写 ^= 32 相当于 -32）
// 大写变小写、小写变小写：字符 |= 32 （大写 |= 32 就相当于+32，小写 |= 32 不变）
// 大写变大写、小写变大写：字符 &= -33 （大写 &= -33 不变，小写 &= -33 相当于 -32）
func toLowerCase2(s string) string {

	lower := &strings.Builder{}
	lower.Grow(len(s))

	for _, ch := range s {
		// 65 A, 90 Z -> 01011010
		// 32 space  00100000
		// 01011010 | 00100000 = 01111010 == z
		if 65 <= ch && ch <= 90 {
			ch |= 32
		}

		lower.WriteRune(ch)
	}

	return lower.String()
}
