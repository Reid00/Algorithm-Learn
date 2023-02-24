package main

// function: heap sort

// []int{4,6,8,5,9}
// len = 5

func HeapSort(a []int) []int {

	// 将无序数组a， 构建成一个大根堆
	for i := len(a)/2 - 1; i >= 0; i-- { // 从最后一个非叶子结点开始，因为叶子节点没有子节点不需要进行堆化[没有左右孩子]
		heapify(a, i, len(a))
	}

	// 交换a[0] 和 a[len(a)-1], 这样最大的元素就在最后面了，然后对除了最后一个元素以外，再找出最大堆
	// 等价于删除堆顶元素 (把堆顶和堆位交换数据 然后保持a[:len(a)-1] 为堆结构）
	for i := len(a) - 1; i >= 1; i-- {
		// 从后往前填充
		swap(a, 0, i) // 第一个元素始终是堆顶
		heapify(a, 0, i)
	}
	return a
}

// 对a[i] 进行上浮/下沉堆化，保持堆结构
func heapify(a []int, i, length int) {
	for {
		// 左孩子节点的索引，从零开始
		l := 2*i + 1
		// 右孩子
		r := 2*i + 2

		// largest 保存根，左，右 三者中的较大值的索引
		largest := i

		// 存在左节点，且左节点值比较大，则取左节点
		if l < length && a[l] > a[largest] {
			largest = l
		}

		// 存在有节点，且值比较大，则取右节点
		if r < length && a[r] > a[largest] {
			largest = r
		}

		// 如果根节点比较大，则不用下沉
		if largest == i {
			break
		}

		// 如果跟节点比较小， 则交换值，并且继续下沉
		swap(a, i, largest)
		// 交换之后，largest 索引代表了原先的父节点a[i]，要保证下沉后 对应的子树同样要保持堆结构，把下沉后的a[i]数据继续堆化
		i = largest
	}

}

// below is declare
func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}