package array

// maxArea  双指针 找出最大的面积
func maxArea(height []int) int {

	l, r := 0, len(height)-1

	maxVal := 0

	for l < r {

		val := min(height[l], height[r]) * (r - l)
		if maxVal < val {
			maxVal = val
		}

		if height[l] < height[r] {
			// 如果左侧指针的长度 小于右侧
			// 左侧指针左移 有可能会加大面积
			l++
		} else {
			r--
		}

	}
	return maxVal
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
