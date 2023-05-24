package stack

// nextGreaterElements2 与 594 不同的是，这是一个循环数组
// 思想上需要O(n) 单调栈  + 循环数组拉平
func nextGreaterElements2(nums []int) []int {

	n := len(nums)

	// 环形数组需要在逻辑层拉平，就是把第一个元素后面的元素放到前面
	ans := make([]int, n)

	for i := range ans {
		ans[i] = -1
	}

	// 存储nums 的index
	stack := make([]int, 0)

	// 环形数组，逻辑层拉平
	// 从左向右遍历，当前元素是栈顶元素的next elem
	for i := 0; i < n*2-1; i++ {
		// i % n 做逻辑拉平
		// 栈顶元素小于当前元素，则当前元素是next greater element
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i%n)

	}

	return ans
}
