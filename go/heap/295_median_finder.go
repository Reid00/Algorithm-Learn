package heap

import "container/heap"

// 数据流的中位数
// 中位数将数组分为两个部分，左边的元素比其小，右边的比其大
// 左边大顶堆，右边小顶堆，小的加左边，大的加右边，
// 平衡俩堆数，新加就弹出，堆顶给对家，奇数取多的，偶数取除2.
type MedianFinder struct {
	qMin bigHeap // 左边的元素，大根堆
	qMax sHeap   // 右边的元素，小根堆
}

func Constructor2() MedianFinder {
	qMin := make(bigHeap, 0)
	qMax := make(sHeap, 0)
	heap.Init(&qMin)
	heap.Init(&qMax)

	return MedianFinder{
		qMin: qMin,
		qMax: qMax,
	}
}

// AddNum 添加一个元素
// 如果总元素个数是奇数，qMin 中的数的数量比 qMax 多一个，此时中位数为 queMin 队头
// 当累计添加的数的数量为偶数时，两个优先队列中的数的数量相同，此时中位数为它们的队头的平均值。
func (m *MedianFinder) AddNum(num int) {
	// 当不存在元素的时候 num 加入到qMin 中
	// 或者挡 num < qMin 中的最大值时，加入左边的堆中，并且将此时元素数量发生变化
	// 将qMin 中的堆顶元素(最大值) 移到右边列表
	if m.qMin.Len() == 0 || num < m.qMin[0] {
		// 入堆
		heap.Push(&m.qMin, num)
		// 堆顶最大值
		// 不能直接pop 因为第一个元素加入一定是qMin 中，但是Pop 会出错
		// b := heap.Pop(&m.qMin).(int)
		// heap.Push(&m.qMax, b)

		// 最多相差一个值
		if m.qMax.Len()+1 < m.qMin.Len() {
			b := heap.Pop(&m.qMin).(int)
			heap.Push(&m.qMax, b)
		}

	} else {
		heap.Push(&m.qMax, num)
		// 堆顶最小值
		if m.qMax.Len() > m.qMin.Len() {
			s := heap.Pop(&m.qMax).(int)
			heap.Push(&m.qMin, s)
		}

	}

}

// FindMedian 找到中位数
// 如果总元素个数是奇数，qMin 中的数的数量比 qMax 多一个，此时中位数为 queMin 队头
// 如果是偶数，(qMin 中的最大值 + qMax 中的最小值) /2
func (m *MedianFinder) FindMedian() float64 {
	// 奇数
	if m.qMin.Len() > m.qMax.Len() {
		return float64(m.qMin[0])
	}

	return float64(m.qMax[0]+m.qMin[0]) / 2
}

// 小根堆
// []int{sum, i,j} => 两个切片的和，对应的下标
type sHeap []int

func (h sHeap) Less(i, j int) bool {
	// > 大根堆， < 小根堆
	return h[i] < h[j]
}

func (h sHeap) Len() int {
	return len(h)
}

// Swap 不会改变切片大小，所以用值接受者和指针接受者，效果相同
func (h *sHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *sHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *sHeap) Pop() interface{} {

	elem := (*h)[len(*h)-1]

	*h = (*h)[:len(*h)-1]
	return elem
}
