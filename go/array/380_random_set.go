package array

import "math/rand"

type RandomizedSet struct {
	data []int
	hmap map[int]int
}

func Constructor2() RandomizedSet {
	return RandomizedSet{
		data: make([]int, 0),
		hmap: make(map[int]int),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	// 如果已经存在返回false
	if _, ok := r.hmap[val]; ok {
		return false
	}
	// 不存在写入集合中
	// 不需要长度-1，新的长度应该就是长度-1+1
	idx := len(r.data)
	r.data = append(r.data, val)
	r.hmap[val] = idx
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	// 如果不存在，返回false
	if _, ok := r.hmap[val]; !ok {
		return false
	}

	// 1. 在map中找到需要删除元素的下标，并删除
	// 2. 在data 中，找到该元素
	// 2.1 如果在data 中是last_index 直接删除
	// 2.2 如果是中间元素，和last 元素交换后，删除最后一个元素保证O(1)的时间复杂度
	idx := r.hmap[val]
	delete(r.hmap, val)

	last_index := len(r.data) - 1
	last := r.data[last_index]
	// 删除 最后的元素
	r.data = r.data[:last_index]

	if idx != last_index {
		r.data[idx] = last
		r.hmap[last] = idx
	}

	return true
}

func (r *RandomizedSet) GetRandom() int {
	randIdx := rand.Intn(len(r.data))
	return r.data[randIdx]
}
