package string

func reverseOnlyLetters(s1 string) string {
	s := []byte(s1)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		for !(s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z') && i < j {
			i++
		}

		for !(s[j] >= 'a' && s[j] <= 'z' || s[j] >= 'A' && s[j] <= 'Z') && i < j {
			j--
		}
		s[i], s[j] = s[j], s[i]
	}

	return string(s)
}
