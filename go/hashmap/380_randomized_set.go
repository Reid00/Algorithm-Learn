package hashmap

import "math/rand"

// O(1) 时间插入、删除和获取随机元素
// O(1)获取随机元素  边长数组可以实现
// O(1) 插入、删除 hashmap 可以实现
type RandomizedSet struct {
	data []int       // 存储元素
	idx  map[int]int // 存储每个元素在切片中的下标 {val： idx}
}

func ConstructorR() RandomizedSet {
	data := make([]int, 0)
	idx := make(map[int]int)
	return RandomizedSet{
		data: data,
		idx:  idx,
	}
}

// Insert 插入一个元素
func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.idx[val]; ok {
		return false
	}
	// 先存储下标
	r.idx[val] = len(r.data)
	// 数据
	r.data = append(r.data, val)
	return true
}

// Remove 删除元素时需要注意O(1)的话 要把需要删除的元素移到最后面
func (r *RandomizedSet) Remove(val int) bool {
	if idx, ok := r.idx[val]; ok {

		// 在data 中的索引下标是 idx
		last := len(r.data) - 1

		// first update idx
		r.idx[r.data[last]] = idx

		// second swap data
		r.data[idx] = r.data[last]

		delete(r.idx, val)
		// third, remove last
		r.data = r.data[:last]
		return true
	}

	return false
}

func (r *RandomizedSet) GetRandom() int {
	// return r.data[rand.Intn(len(r.data))]
	randIdx := rand.Intn(len(r.data))
	return r.data[randIdx]
}
