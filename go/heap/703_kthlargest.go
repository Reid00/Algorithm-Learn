package heap

// KthLargest 第K大元素，用长度为K 的最小堆
type KthLargest struct {
	// 存放堆数据
	data []int
	// 堆的最大长度，默认等于k
	size int
	// 堆中已有的数据
	count int
}

func Constructor(k int, nums []int) KthLargest {

	var kth KthLargest
	kth.size = k
	for _, v := range nums {
		kth.Add(v)
	}
	return kth
}

// Add 添加元素，并保证add 之后保证符合堆的特征
func (k *KthLargest) Add(val int) int {

	if k.count < k.size-1 {
		k.data = append(k.data, val)
		k.count++
	} else if k.count == k.size-1 {
		// 容量刚好满，数据进行堆化
		k.data = append(k.data, val)
		k.count++

		// 从第一个非叶子节点 (len(k.data)/2-1) 开始堆化
		for i := len(k.data)/2 - 1; i >= 0; i-- {
			k.heapify(i, len(k.data))
		}
	} else {
		// 容量已满的情况下，添加后数据后需要删除一个维持长度k
		if val > k.data[0] {
			// val 大于堆顶元素，把val 直接替换堆顶，再一次堆化
			k.data[0] = val
			k.heapify(0, len(k.data))
		}
	}
	return k.data[0]
}

// heapify 将k 中data 数组里面的idx 数据进行堆化
// length 表示data 数组的有效长度
func (k *KthLargest) heapify(idx, length int) {

	for {
		// 左右孩子
		l := 2*idx + 1
		r := 2*idx + 2

		// 因为是小根堆 所以 找出idx 左右孩子中的最小值
		var smallest = idx

		if l < length && k.data[smallest] > k.data[l] {
			smallest = l
		}

		if r < length && k.data[smallest] > k.data[r] {
			smallest = r
		}
		// 目前已经符合小根堆的特征了，父节点的值小于等于左右孩子
		if smallest == idx {
			break
		}

		// 否则 交换idx 和  smallest 所在的元素
		k.data[idx], k.data[smallest] = k.data[smallest], k.data[idx]

		// 交换之后 idx 作为父节点已经满足堆的特征，但是smallest 作为父节点同样保证堆化
		idx = smallest
	}
}
