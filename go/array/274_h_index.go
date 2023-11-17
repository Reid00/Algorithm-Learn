package array

// hIndex 具体定义可以Google 下，重点理解其概念
// 如果有N篇论文至少被引用了N次，那么这个h-index 是N
// 1. 将其发表的所有SCI论文按被引次数从高到低排序；
// 2. 从前往后查找排序后的列表，直到某篇论文的序号大于该论文被引次数。所得序号减一即为H指数
func hIndex(citations []int) int {
	// 二分
	l, r := 0, len(citations)
	// 引用次数
	cnt := 0

	for l < r {
		// 二分从中间开始查找
		mid := l + (r-l)>>1

		for _, v := range citations {
			// 引用次数v 至少mid 次
			if mid <= v {
				cnt++
			}
		}

		if cnt >= mid {
			// 要找的答案在 [mid,right] 区间内
			l = mid
		} else {
			// 要找的答案在 [0,mid) 区间内
			r = mid - 1
		}
	}
	return l
}
