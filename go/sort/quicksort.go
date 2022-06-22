package main

// function: QuickSort

func quickSort(a []int) []int {
	_quickSort(a, 0, len(a)-1)
	return a
}

// 原地交换，所以传入索引
func _quickSort(nums []int, start, end int) {
	if start < end {
		// 分治法
		pivot := partition(nums, start, end)
		_quickSort(nums, start, pivot-1)
		_quickSort(nums, pivot+1, end)
	}
}

// 对nums[start:end+1] 进行分区，找到一个pivot 使得左边的比它小，右边的比他大
func partition(nums []int, start, end int) int {

	p := nums[end] // 需要比较的元素
	s := start     // 支点的位置，s 左侧比p 小，右侧比p 大

	for i := start; i < end; i++ {
		if nums[i] < p {
			swap(nums, s, i)
			s++
		}
	}
	// 循环结束之后，交换p 与s 的位置
	swap(nums, s, end)
	return s
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
