package heap

import "container/heap"

// findKthLargest 找到数组中第k大的元素, Miniheap
func findKthLargest(nums []int, k int) int {
	// heapInt define in 40_offer_get_leatest_nums

	h := make(heapInt, 0)

	// 构建长度为k 的小根堆
	heap.Init(&h)

	for _, v := range nums {
		heap.Push(&h, v)
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	return h[0]
}
