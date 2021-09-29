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
		_quickSort(nums, 0, pivot-1)
		_quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start // 记录最终基准元素的index

	for j := start; j < end; j++ {
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, end) // 循环结束后，i 是第一个比基准元素大的值，将它跟基准元素进行交换
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
