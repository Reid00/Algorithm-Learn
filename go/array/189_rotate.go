package array

// 多次数组翻转
// [1,2,3,4,5,6]  k =2  ret = [5,6,1,2,3,4]
// 1. 全部翻转 ->  [6,5,4,3,2,1]
// 2. 前k 个翻转 -> [5,6,4,3,2,1]
// 2. 后n-k 个翻转 -> [5,6,1,2,3,4]
// 需要注意的是k 可能大于n=len(nums)， 那么k = k %n

func rotate(nums []int, k int) {

	n := len(nums)
	k = k % n
	reverse2(nums, 0, n-1)
	reverse2(nums, 0, k-1)
	reverse2(nums, k, n-1)

}

func reverse2(nums []int, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
