package binarysearch

func findMin2(nums []int) int {

	n := len(nums)

	l, r := 0, n-1

	for l < r {

		mid := l + (r-l)>>1

		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			// nums[mid] = nums[r]
			// NOTE this is the key
			r--
		}

	}
	return nums[l]
}
