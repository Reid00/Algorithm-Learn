package main

func mergeSort(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}

	mid := length / 2

	left := mergeSort(a[:mid])
	right := mergeSort(a[mid:])

	return merge(left, right)
}

func merge(left, right []int) (result []int) {

	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] { // left 和 right 中的元素，谁小，先合并谁
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	// 合并完，合并两个切片中剩余的元素
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
