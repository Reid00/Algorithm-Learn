package sort

import "testing"

var nums = []int{1, 5, 13, 6, 2}

func TestQuickSort(t *testing.T) {
	QuickSort(nums)
	t.Log(nums)
}

func TestBubleSort(t *testing.T) {
	bubbleSort(nums)
	t.Log(nums)
}

func TestMergeSort(t *testing.T) {
	ret := MergeSort(nums)
	t.Log(ret)
}

func TestSelectionSort(t *testing.T) {
	selectionSort(nums)
	t.Log(nums)
}

func TestHeapSort(t *testing.T) {
	heapSort(nums)
	t.Log(nums)
}
