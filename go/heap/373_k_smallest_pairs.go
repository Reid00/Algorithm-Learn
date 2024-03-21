package heap

import "container/heap"

// kSmallestPairs 查找和最小的 K 对数字
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	// 将 {sum, i, j} 三元组以sum 为形势压入小根堆中，堆中k 个元素
	h := make(smallHeap, 0)

	ans := make([][]int, 0)
	// 将nums1,nums2 拼成堆内元素的形式，需要注意的是可以只需要k个元素即可
	// 另外拼接的时候可以把较短的数组放到前面
	n, m := len(nums1), len(nums2)
	flag := n > m
	// 如果n 较长，将nums2 放到前面，需要注意的是ans 里面第一个元素是要求来自nums1 的，append 的时候需要注意
	if flag {
		n, m, nums1, nums2 = m, n, nums2, nums1
	}

	// 最小堆中放k个元素足够了
	if n > k {
		n = k
	}

	// 添加到h 中
	for i := 0; i < n; i++ {
		h = append(h, []int{nums1[i] + nums2[0], i, 0})
	}

	// 堆化 heapify
	heap.Init(&h)
	for h.Len() > 0 && len(ans) < k {
		// 堆顶是最小值
		small := heap.Pop(&h).([]int)

		x, y := small[1], small[2]
		// 是否交换过nums1 和 nums2
		// flag true 交换过，nums2 是原始的nums1
		if flag {
			ans = append(ans, []int{nums2[y], nums1[x]})
		} else {
			ans = append(ans, []int{nums1[x], nums2[y]})
		}

		// 入堆 下一个元素
		if y < m-1 {
			heap.Push(&h, []int{nums1[x] + nums2[y+1], x, y + 1})
		}
	}
	return ans
}

// 小根堆
// []int{sum, i,j} => 两个切片的和，对应的下标
type smallHeap [][]int

func (h smallHeap) Less(i, j int) bool {
	// > 大根堆， < 小根堆
	return h[i][0] < h[j][0]
}

func (h smallHeap) Len() int {
	return len(h)
}

// Swap 不会改变切片大小，所以用值接受者和指针接受者，效果相同
func (h *smallHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *smallHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *smallHeap) Pop() interface{} {

	elem := (*h)[len(*h)-1]

	*h = (*h)[:len(*h)-1]
	return elem
}
