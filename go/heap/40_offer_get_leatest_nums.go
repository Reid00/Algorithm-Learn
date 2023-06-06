package heap

import "container/heap"

// getLeastNumbers 输入整数数组 arr ，找出其中最小的 k 个数。
// 例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}

	// 构建小根堆
	var h heapInt = make([]int, 0)
	heap.Init(&h)

	for _, v := range arr {
		heap.Push(&h, v)
	}

	var ans []int
	for i := 0; i < k; i++ {
		s := heap.Pop(&h)
		ans = append(ans, s.(int))
	}
	return ans
}

// 小根堆
type heapInt []int

func (h heapInt) Less(i, j int) bool {
	// > 大根堆， < 小根堆
	return h[i] < h[j]
}

func (h heapInt) Len() int {
	return len(h)
}

// Swap 不会改变切片大小，所以用值接受者和指针接受者，效果相同
func (h *heapInt) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *heapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heapInt) Pop() interface{} {

	elem := (*h)[len(*h)-1]

	*h = (*h)[:len(*h)-1]
	return elem
}

// 查看堆顶元素
func (h heapInt) Peek() interface{} {
	return h[0]
}
