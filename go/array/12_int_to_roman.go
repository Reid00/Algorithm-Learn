package array

import "strings"

// intToRoman 整数转罗马数字
func intToRoman(num int) string {

	result := []string{}

	for num > 0 {
		if num >= 1000 {
			num -= 1000
			result = append(result, "M")
		} else if num >= 900 {
			num -= 900
			result = append(result, "CM")
		} else if num >= 500 {
			num -= 500
			result = append(result, "D")
		} else if num >= 400 {
			num -= 400
			result = append(result, "CD")
		} else if num >= 100 {
			num -= 100
			result = append(result, "C")
		} else if num >= 90 {
			num -= 90
			result = append(result, "XC")
		} else if num >= 50 {
			num -= 50
			result = append(result, "L")
		} else if num >= 40 {
			num -= 40
			result = append(result, "XL")
		} else if num >= 10 {
			num -= 10
			result = append(result, "X")
		} else if num >= 9 {
			num -= 9
			result = append(result, "IX")
		} else if num >= 5 {
			num -= 5
			result = append(result, "V")
		} else if num >= 4 {
			num -= 4
			result = append(result, "IV")
		} else if num >= 1 {
			num -= 1
			result = append(result, "I")
		}
	}

	ret := strings.Join(result, "")

	return ret
}
