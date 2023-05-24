package stack

// largestRectangleArea 柱状图中最大的矩形柱状图中最大的矩形
// 依次遍历柱形的高度，对于每一个高度分别向两边扩散，求出以当前高度为矩形的最大宽度多少。
func largestRectangleArea(heights []int) int {
	// 单调栈，把高度从低到高排序

	n := len(heights)
	stack := make([][]int, 0)

	ans := 0

	// 这里有个小技巧，可以在给定的数组后面再加一个高度为 0 的柱子
	// 在不影响结果的情况下，可以直接将栈清空
	for i := 0; i <= n; i++ {
		// 遍历所有的高度
		height := 0
		if i < n {
			height = heights[i]
		}

		if len(stack) == 0 {
			// 记录矩形的宽和高
			stack = append(stack, []int{1, height})
			continue
		}

		if height == stack[len(stack)-1][1] {
			// 矩形的宽 加一
			stack[len(stack)-1][0] = stack[len(stack)-1][0] + 1
			continue
		}

		if height > stack[len(stack)-1][1] {
			stack = append(stack, []int{1, height})
			continue
		}

		width := 0
		for len(stack) > 0 && height < stack[len(stack)-1][1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width += top[0]
			ans = max(ans, width*top[1])
		}

		stack = append(stack, []int{width + 1, height})

	}
	return ans
}

// 暴力解法, timeout 未通过
func largestRectangleArea2(heights []int) int {
	// 固定高度，枚举宽度

	n := len(heights)
	ans := 0

	for i := 0; i < n; i++ {
		left, right := i, i
		// 以i 为中心，向左右枚举，获取宽度可以为heights[i] 的最大宽度

		for left-1 >= 0 && heights[left-1] >= heights[i] {
			left--
		}

		for right+1 < n && heights[right+1] >= heights[i] {
			right++
		}

		ans = max(ans, (right-left+1)*heights[i])
	}
	return ans
}
