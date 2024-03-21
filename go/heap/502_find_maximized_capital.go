package heap

import (
	"container/heap"
	"slices"
)

// findMaximizedCapital IPO
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// 根据贪心的策略，我们要想总的资本最大，那么每一次都要在能启动的项目中选择利润最大的那个。
	// 因此我们可以使用一个大顶堆来维护当前所有能启动的项目，堆顶项目即为利润最大的那个。

	// 对利润维护大根堆
	h := make(bigHeap, 0)

	pairs := make([]Pair, len(profits))
	for i, p := range profits {
		pairs[i] = Pair{
			p: p,
			c: capital[i],
		}
	}
	// 对启动资本升序排列
	slices.SortFunc(pairs, func(i, j Pair) int { return i.c - j.c })

	// i 从0 开始启动资本是最小的
	for i := 0; k > 0; k-- {
		// 贪心: 每次选择的全部成本要小于等于资金w 的，然后拿出利润最大的
		for i < len(capital) && pairs[i].c <= w {
			// 将能选择的项目利润压入堆中，堆顶是利润最大的
			heap.Push(&h, pairs[i].p)
			i++
		}

		// 防止没有可以选择的项目
		if h.Len() == 0 {
			break
		}
		// 资本在累积
		w += heap.Pop(&h).(int)
	}

	return w
}

// Pair 利润和资本的pair对
// 大根堆根据利润最大的进行排序
type Pair struct {
	p int
	c int
}

// 大根堆
type bigHeap []int

func (h bigHeap) Less(i, j int) bool {
	// > 大根堆， < 小根堆
	return h[i] > h[j]
}

func (h bigHeap) Len() int {
	return len(h)
}

// Swap 不会改变切片大小，所以用值接受者和指针接受者，效果相同
func (h *bigHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *bigHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *bigHeap) Pop() interface{} {

	elem := (*h)[len(*h)-1]

	*h = (*h)[:len(*h)-1]
	return elem
}

// 查看堆顶元素
func (h bigHeap) Peek() interface{} {
	return h[0]
}
