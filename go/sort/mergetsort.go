package sort

// 归并排序（Merge sort）是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。

// 作为一种典型的分而治之思想的算法应用，归并排序的实现由两种方法：
// 自上而下的递归（所有递归的方法都可以用迭代重写，所以就有了第 2 种方法）；
// 自下而上的迭代；

// 递归的实现
func MergeSort(a []int) []int {
	length := len(a)
	if length < 2 {
		return a
	}

	mid := length / 2

	left := MergeSort(a[:mid])
	right := MergeSort(a[mid:])

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
