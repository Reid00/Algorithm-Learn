package main

// 1. 从数列中挑出一个元素，称为 "基准"（pivot）;
// 2. 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面
// （相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
// 3. 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序；

func quickSort(a []int) {
	_quickSort(a, 0, len(a)-1)
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

// 对nums[start:end+1] 进行分区, nums[end] 默认为pivot
func partition(nums []int, start, end int) int {

	p := nums[end] // 需要比较的元素，基准pivot
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
