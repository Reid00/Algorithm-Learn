package binarysearch

import "sort"

// firstBadVersion 第一个错误的版本
func firstBadVersion(n int) int {

	left := 1
	right := n

	var isBadVersion func(int) bool

	for left < right {
		mid := left + (right-left)>>1
		if isBadVersion(mid) {
			// mid 右侧都是bad version 第一个badversion 在mid 左侧
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func firstBadVersion2(n int) int {
	var isBadVersion func(int) bool
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })

}
