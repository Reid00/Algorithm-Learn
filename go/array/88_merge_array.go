package main

// merge O(n) 的时间复杂度
// O(m+n)的空间复杂度
func merge(nums1 []int, m int, nums2 []int, n int) {

	sorted := make([]int, 0, m+n)

	p1, p2 := 0, 0

	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}

		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}

		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
		}

	}
	copy(nums1, sorted)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {

	p1, p2 := m-1, n-1
	idx := m + n - 1

	for p1 > 0 && p2 > 0 {
		if nums1[p1] > nums2[p2] {
			nums1[idx] = nums1[p1]
			p1--
		} else {
			nums1[idx] = nums2[p2]
			p2--
		}
		idx--
	}

	for p2 > 0 {
		nums1[idx] = nums2[p2]
		p2--
		idx --
	}
}
