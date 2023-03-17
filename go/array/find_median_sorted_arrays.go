package array

// 暴力遍历合并数组
// O(m+n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// 长度m, n  的数组中位数为:
	// 1. m + n 为奇数: nums[(m+n)/2]
	// 2. m + n 为偶数: (nums[(m+n)/2-1] + nums[(m+n)/2] )/2

	// 遍历两个数组 并排序合并

	var i, j int
	var merged []int

	m, n := len(nums1), len(nums2)

	for i < m && j < n {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}

	// 处理剩下未被merged 的元素
	if i < m {
		merged = append(merged, nums1[i:]...)
	}
	if j < n {
		merged = append(merged, nums2[j:]...)
	}

	var median float64
	if (m+n)&1 == 0 {
		// even num
		l := (m+n)/2 - 1
		r := l + 1
		median = float64(merged[l]+merged[r]) / 2
	} else {
		// odd num
		median = float64(merged[(m+n)/2])
	}
	return median
}

// 二分法
// 算法的时间复杂度应该为 O(log (m+n))
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	// 分别找到两个数组所在k位(中位数)数字
	// if m + n是奇数，直接是 第 （ m +n ) /2
	// even: (m+n)/2 -1 和 （m+n)/2 的平均值

	if l := len(nums1) + len(nums2); l&1 == 0 { // even num
		k1 := getKth(nums1, nums2, l/2 -1)
		k2 := getKth(nums1, nums2, l/2)
		return float64(k1+k2) / 2
	} else {
		return float64(getKth(nums1, nums2, l/2))
	}

}

// 返回两个有序数组中，第k小的元素 k >= 0
// 二分法
func getKth(nums1, nums2 []int, k int) int {

	// 第k 小的元素 存在于，nums1和nums2 左半区, 或者右半区
	// 利用二分法，一次性排除一半，O(log (m +n))

	for {
		l1, l2 := len(nums1), len(nums2)
		mid1, mid2 := l1/2, l2/2

		if l1 == 0 {
			return nums2[k]
		}

		if l2 == 0 {
			return nums1[k]
		}

		if k == 0 {
			return min(nums1[0], nums2[0])
		}

		if k <= mid1+mid2 {
			// 第k个元素在nums1 nums2 的左半部，
			// 删除右半部的元素
			if nums1[mid1] < nums2[mid2] {
				//因为k 在nums1 nums2 的前半部分
				// 并且nums2 的后半部分较大，需要排除掉
				nums2 = nums2[:mid2]
			} else {
				nums1 = nums1[:mid1]
			}
		} else {
			// 如果k在nums1 nums2的右半部
			if nums1[mid1] < nums2[mid2] {
				// 排除nums1 的前半部分
				nums1 = nums1[mid1+1:]
				// 删除了前面的mid1 的数据，需要更新k的索引
				// k 是从0 开始的，要 + 1
				k -= mid1 + 1
			} else {
				nums2 = nums2[mid2+1:]
				k -= mid2 + 1
			}
		}

	}

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
