package array

//  canJump 跳跃游戏
func canJump(nums []int) bool {

	// 贪心算法，每次都跳跃到最大位置，更新maxPos, 观察是否能到达len(nums)-1
	// 注意: 遍历到下标位置时，先考虑能否到达该位置，再去更新上面的maxPos

	maxPos := 0

	for i, v := range nums {
		// 判断是否能到i 这个位置
		if i <= maxPos {
			if maxPos < i + v {
				maxPos = i + v
			}
		}
	}
	return maxPos >= len(nums)-1
}
