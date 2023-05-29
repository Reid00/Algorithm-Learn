package queue

// maxSlidingWindow 滑动窗口中的最大值
// 单调队列解法
// 保证在window k 中的数据是单调递减的，队列的首部永远是最大值
// window 里面可以选择放数据，或者在nums 中的下标(更简单)
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)

	deque := make([]int, 0)
	ans := make([]int, 0, n-k+1)

	for i, v := range nums {
		// 将出界的下删除, win中的长度超过k
		if i >= k && i-deque[0] >= k {
			deque = deque[1:]
		}

		// 将i添加到队列中，并保证队列是递减的
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= v {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i >= k-1 {
			ans = append(ans, nums[deque[0]])
		}
	}

	return ans
}
