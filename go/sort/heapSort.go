package main

// 堆是具有以下性质的完全二叉树：每个结点的值都大于或等于其左右孩子结点的值，称为大顶堆；
// 或者每个结点的值都小于或等于其左右孩子结点的值，称为小顶堆。

// 该数组从逻辑上讲就是一个堆结构，我们用简单的公式来描述一下堆的定义就是：
// 大顶堆：arr[i] >= arr[2i+1] && arr[i] >= arr[2i+2]
// 小顶堆：arr[i] <= arr[2i+1] && arr[i] <= arr[2i+2]

func heapSort(nums []int) {
	buildHeap(nums, 0, len(nums))

	// 将最大值nums[0] 和最后一个元素交换，剩下的nums[:n-1] 继续堆化
	for i := len(nums) - 1; i >= 1; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapfy(nums, 0, i)
	}
}

// 构建出堆
// 堆的构建过程就是从最后一个非叶子结点(n/2 -1)开始从下往上调整。
func buildHeap(nums []int, i, length int) {

	for i := len(nums)/2 - 1; i > 0; i-- {
		heapfy(nums, i, length)
	}
}

// 堆化，对i 所在的位置的节点进行堆化，保证两点
// 1. i 所在的树 满足堆的特性
// 2. i 堆化后，如果有元素交换，保证交换后 子树也满足堆的特性
func heapfy(nums []int, i, length int) {

	for {
		// 左右子树的index
		l, r := 2*i+1, 2*i+2
		// i 所在二叉树的最大元素的索引
		largest := i

		if l < length && nums[l] > nums[largest] {
			largest = l
		}

		if r < length && nums[r] > nums[largest] {
			largest = r
		}

		// i 就是最大值 无需交换元素
		if largest == i {
			break
		}

		// i 不是最大值，则进行堆化，保证堆的特征
		nums[i], nums[largest] = nums[largest], nums[i]

		// 保证i 子树 满足堆的特征,
		// largest 是i 子树的索引
		i = largest
	}
}
