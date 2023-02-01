package main

// 1. 在未排序的数组中，先找出最大的元素
// 2. 在剩余的元素中找出 最大的元素
// 3. 重复1，2

func selectionSort(nums []int) {

	length := len(nums)

	for i := 0; i < length-1; i++ {
		minIndex := i // 最小元素的索引, 初始设为i

		for j := i + 1; j < length; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		// 一轮循环结束后，交换最大元素的位置
		nums[minIndex], nums[i] = nums[i], nums[minIndex]
	}

}
