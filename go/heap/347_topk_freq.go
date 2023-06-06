package heap

import "container/heap"

// topKFrequent 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。
//  用小根堆 解决TopK 大的问题。逻辑是长度为k的堆中，当容量满了之后，需要删除堆顶的最小的元素
// 用hashMap{key: cnt} 表示一个堆的元素
func topKFrequent(nums []int, k int) []int {

	hmap := make(map[int]int)

	for _, v := range nums {
		hmap[v] += 1
	}

	var kv = make(KVList, 0)
	heap.Init(&kv)

	for key, v := range hmap {
		heap.Push(&kv, [2]int{key, v})
		if kv.Len() > k {
			heap.Pop(&kv)
		}
	}

	ans := make([]int, 0)

	for i := 0; i < k; i++ {
		item := heap.Pop(&kv)
		ans = append(ans, item.([2]int)[0])
	}
	return ans
}

// type KV struct {
// 	k int
// 	v int
// }

type KVList [][2]int // []{[k, cnt], [k, cnt]...}

func (kv KVList) Len() int {
	return len(kv)
}

func (kv KVList) Less(i, j int) bool {
	return kv[i][1] < kv[j][1]
}

func (kv KVList) Swap(i, j int) {
	kv[i], kv[j] = kv[j], kv[i]
}

func (kv *KVList) Push(x interface{}) {
	(*kv) = append((*kv), x.([2]int))
}

func (kv *KVList) Pop() interface{} {
	item := (*kv)[len(*kv)-1]
	(*kv) = (*kv)[:len(*kv)-1]

	return item
}
