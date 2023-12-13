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

// threeSum 三数之和
func threeSum2(nums []int) [][]int {

	ans := make([][]int, 0)

	// 先排序，固定一个数字之后，再用左右指针去找组合
	// 注意1: 答案中不可以包含重复的三元组，则固定一个元素后，第二个第三个元素应该过滤掉重复的元素
	sort.Ints(nums)

	for first := 0; first < len(nums); first++ {

		// 优化1: 第一个元素大于0 了，则后续升序排列 不存在pairs
		if nums[first] > 0 {
			break
		}

		// 优化3:
		if first+2 < len(nums) && nums[first]+nums[first+1]+nums[first+2] > 0 {
			break
		}

		// 避免注意1 中的问题， 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := len(nums) - 1
		// 枚举第二个和第三个元素
		for second := first + 1; second < len(nums); second++ {

			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 优化2: 如果三数之和小于0 则直接增加second, 进行下一次循环，避免无效遍历
			if nums[first]+nums[second]+nums[third] < 0 {
				continue
			}

			// 枚举第三个元素(在第二个循环之内)，此种办法每次枚举第三个元素，
			// 都会从头遍历，增加了时间复杂度
			// 把third 先确定之后，在枚举第二个元素，如果元素之和都大于0 因为有序排列，则此时没有结果
			// third := len(nums) - 1
			for second < third && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}

			// 以first 开始的三元组没有可以满足的
			if second == third {
				break
			}

			if nums[first]+nums[second]+nums[third] == 0 {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}

		}

	}

	return ans
}
