package array

// twoSum 两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	// 仅存在一个有效答案，非递减顺序排列 (不可以 重复使用相同的元素)
	// O(1) 的额外空间复杂度

	// 左右指针
	l, r := 0, len(numbers)-1

	for l < r {

		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}

		if numbers[l]+numbers[r] < target {
			l++
			r = len(numbers) - 1
		}

		if numbers[l]+numbers[r] > target {
			r--
		}
	}

	return []int{-1, -1}
}
