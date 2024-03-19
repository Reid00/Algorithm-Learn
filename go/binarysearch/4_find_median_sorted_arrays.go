package binarysearch

// findMedianSortedArrays 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// 偶数
	if total := len(nums1) + len(nums2); total&1 == 0 {
		e1 := getKNums(nums1, nums2, total/2)
		e2 := getKNums(nums1, nums2, total/2+1)
		return float64(e1+e2) / 2
	} else {
		return float64(getKNums(nums1, nums2, total/2))
	}

}

// getKNums 在有序数组nums1/nums2 中获取第k 小的元素(下标从0开始)
func getKNums(nums1, nums2 []int, k int) int {
	// log(m+n) 的复杂度，二分
	// 每次抛弃nums1 nums2 中一般的元素

	for {

		n1, n2 := len(nums1), len(nums2)
		mid1, mid2 := n1>>1, n2>>1

		if n1 == 0 {
			return nums2[k]
		}

		if n2 == 0 {
			return nums1[k]
		}

		// 循环终止的条件, 第0小的元素
		// 每次抛弃一半的元素之后，如果是前半部分的元素，要更新k 的值，会变小
		if k == 0 {
			return min(nums1[0], nums2[0])
		}

		// k在 两个数组的前半部分较小的元素中， 要舍弃后半部分较大的
		// 特别分析，k == mid1 + mid2 的时候，实际上合并有mid1 + 1 + mid2 + 1 个元素
		// 而 第k 小，实际上是 k + 1个元素，因此抛弃元素的时候，mid1/mid2 是可以不要的
		if k <= mid1+mid2 {

			// 说明nums2 的后半部分元素较大，舍弃
			if nums1[mid1] < nums2[mid2] {
				nums2 = nums2[:mid2]
			} else {
				nums1 = nums1[:mid1]
			}

		} else {
			// k 元素在 两个数组中较大的后半部分，要舍弃前半部分较小的
			// 说明 nums1 前半部分较小
			if nums1[mid1] < nums2[mid2] {
				nums1 = nums1[mid1+1:]
				// 更新k 的值, 删除(mid1+1) 个元素，因为mid1 的下标从0开始的
				k -= mid1 + 1
			} else {
				nums2 = nums2[mid2+1:]
				k -= mid2 + 1
			}

		}

	}

}
