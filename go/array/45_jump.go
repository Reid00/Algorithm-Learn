package array

// jump 跳跃游戏2
func jump(nums []int) int {

	n := len(nums)

	steps := 0        // 跳跃的步数
	farthest := 0     // 本次跳跃的最远位置
	nextFarthest := 0 // 下次跳跃的最远位置

	// 不需要在最后的位置n-1 进行尝试跳跃，因为已经在最后的位置了
	for i := 0; i < n-1; i++ {
		nextFarthest = max(nextFarthest, i+nums[i])

		// 本次跳跃结束
		if i == farthest {
			// 更新下一次跳跃的最大位置
			farthest = nextFarthest
			steps++
		}
	}
	return steps
}
