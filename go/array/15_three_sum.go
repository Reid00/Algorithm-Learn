package array

import "sort"

func threeSum(nums []int) [][]int {

	// 三数之和 a + b + c = 0
	// 先排序，再盯住first element 的情况下用左右指针找到second + third = -first

	// 去除特殊情况:
	sort.Ints(nums)
	if len(nums) < 3 {
		return nil
	}

	ans := make([][]int, 0)

	n := len(nums)

	// first 要< n-1 second 才不会out of bound
	for first := 0; first < n; first++ {
		// 固定第一个元素，然后枚举第二个元素, 第三个元素

		// 避免枚举重复的元素, n > 0 第二个元素开始比较
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		target := -1 * nums[first]
		third := n - 1

		for second := first + 1; second < n; second++ {

			// 排除 和之前重复的元素避免形成重复的结果
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 如果 nums[second] + nums[third] >target 说明 third 太大，third --
			// 同时 second < third 否则 third 就不存在了
			for second < third && nums[second]+nums[third] > target {
				third--
			}

			// 本轮循环， first 固定的情况下没有找到合适的值
			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}

	}

	return ans
}
