package array

// romanToInt 罗马数字转整数
func romanToInt(s string) int {

	hmap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// special := []string{
	// 	//4,    9	 40,   90,   400,  900
	// 	"IV", "IX", "XL", "XC", "CD", "CM",
	// }
	// special := []byte{
	// 	//4,    9	 40,   90,   400,  900
	// 	'I', 'X', 'C',
	// }

	result := 0
	i := 0

	for i < len(s)-1 {
		tmp := 0
		ch := s[i]
		if ch == 'I' && s[i+1] == 'V' {
			tmp = 4
			i += 2
		} else if ch == 'I' && s[i+1] == 'X' {
			tmp = 9
			i += 2
		} else if ch == 'X' && s[i+1] == 'L' {
			tmp = 40
			i += 2
		} else if ch == 'X' && s[i+1] == 'C' {
			tmp = 90
			i += 2
		} else if ch == 'C' && s[i+1] == 'D' {
			tmp = 400
			i += 2
		} else if ch == 'C' && s[i+1] == 'M' {
			tmp = 900
			i += 2
		} else {
			tmp = hmap[byte(ch)]
			i += 1
		}
		result += tmp
	}

	// 倒数第二个是特殊字符的话，i+2 会超过 字符串下标
	if i <= len(s)-1 {
		last := hmap[s[len(s)-1]]
		result += last
	}

	return result
}
