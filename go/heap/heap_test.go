package heap
import (
	"container/heap"
	"fmt"
)

// function: 堆相关的操作

func main() {

	// 测试最大堆

	h := make(MaxHeap, 0)
	heap.Init(&h)

	heap.Push(&h, 8)
	heap.Push(&h, 1)
	heap.Push(&h, 4)
	heap.Push(&h, 5)
	heap.Push(&h, 2)

	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))

}

// Method 2: 用标准库中的"heap" package
// 定义一个最大堆，并实现heap.Interface 接口
type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
	old := *m
	res := old[len(old)-1]
	*m = old[:len(old)-1]
	return res
}

// Method 1: 不用"container/heap"包
type MinHeap struct {
	Element []int
}

// 构造函数
func NewMinHeap() *MinHeap {
	return &MinHeap{
		Element: make([]int, 0),
	}
}

// heap 的主要两种操作: 插入 + 删除
// 插入
func (m *MinHeap) Insert(n int) {
	m.Element = append(m.Element, n) // 先把元素添加到尾部，在进行向上堆化

	// 上浮堆化
	// i := len(m.Element) - 1
	// for ; m.Element[i/2] > n; i /= 2 { // 与父节点比较。i 的父节点为i-1/2, 左孩子2*i + 1, 右 2 * i +2
	// 	m.Element[i] = m.Element[i/2]
	// }
	// 上浮的另一种方式, 通过比较找出n 最终的位置
	j := len(m.Element) - 1 // j 代表最后一个元素
	for {
		i := (j - 1) / 2                           // parent
		if j == i || m.Element[j] > m.Element[i] { //1. i == j 代表在堆顶，0 2. m.Element[j] > m.Element[i] 代表，j 所在位置元素大于父节点，满足最小堆的条件
			break
		}
		// 父子节点交换
		m.Element[j], m.Element[i] = m.Element[i], m.Element[j]

		j = i
	}

	// m.Element[j] = n
}

// 删除并返回最小值
// 将堆树的最后的节点提到根结点，然后删除最小值，然后再把新的根节点放到合适的位置。;最坏时间复杂度同样为O(log N)
func (m *MinHeap) DeleteMin() (int, error) {
	if len(m.Element) <= 1 {
		return 0, fmt.Errorf("MinHeap is an empty")
	}

	min := m.Element[0]
	last := m.Element[len(m.Element)-1]
	m.Element[0] = last
	i := 0 //堆顶元素，最小值的索引
	for {
		left := 2*i + 1
		flag := left                            // flag 比较左右孩子后，去除较小的索引，然后拿到这个索引与父节点交换值
		if left >= len(m.Element) || left < 0 { //只有当前一个元素
			break
		}
		right := 2*i + 2

		if right <= len(m.Element) && m.Element[left] > m.Element[right] { //left 是较小的元素索引, 把left 推到堆顶
			flag = right
		}

		// min下沉, 与较小的孩子节点交换
		m.Element[flag], m.Element[i] = m.Element[i], m.Element[flag]
		i = flag
	}
	m.Element = m.Element[:len(m.Element)-1]
	return min, nil
}
