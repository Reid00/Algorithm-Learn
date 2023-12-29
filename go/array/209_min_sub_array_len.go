package array

import (
	"math"
)

// minSubArrayLen 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	// 快慢指针，弄个滑动窗口
	// 滑动窗口中sum 是两个左右指针夫妻共同的财产。好不容易共同财产大过target，记录一下两指针之间的距离，
	// 当财产满足时，左指针就会胡乱花销(做指针右移动)
	// 为了保证财产不缩水，右指针就会继续工作(右指针右移)， 直到左右指针离得最近的时候

	slow, fast := 0, 0
	// 长度
	ans := math.MaxInt32

	sum := 0

	for fast < len(nums) {

		sum += nums[fast]
		// 财产满足之后，左指针右移，直到不满足大于target
		for sum >= target {
			ans = min(ans, fast-slow+1)
			sum -= nums[slow]
			slow++
		}
		fast++
	}
	if ans == math.MaxInt32 {
		return 0
	}

	return ans
}
