package array

import "math/rand"

type Solution struct {
	data     []int
	original []int
}

func Constructor(nums []int) Solution {
	ori := append([]int(nil), nums...)
	return Solution{
		data:     nums,
		original: ori,
	}
}

func (s *Solution) Reset() []int {

	copy(s.data, s.original)
	return s.data
}

func (s *Solution) Shuffle() []int {
	// 洗牌算法，保证每次取到后的数据都不能再被拿到

	n := len(s.data)
	for i := range s.data {
		// 保证取到的数据 在上次没有取到的数据之后
		j := i + rand.Intn(n-i)
		// 每次随机取之后把取到的数据和 上次的数据换一下
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}

	return s.data
}
