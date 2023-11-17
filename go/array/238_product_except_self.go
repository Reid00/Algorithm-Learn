package array

// productExceptSelf 除自身以外数组的乘积
// O(n) 的空间复杂度
func productExceptSelf(nums []int) []int {

	length := len(nums)
	// 结果集
	result := make([]int, length)
	// l[i] 表示i 左侧所有数之积
	l := make([]int, length)
	l[0] = 1 // l[0]左侧没有数字，但是l[1]要和l[0] * nums[0] 相乘所以保证为nums[0]结果，l[0]=1
	// r[i] 表示i 右侧所有数之积
	r := make([]int, length)
	r[length-1] = 1

	// 填充数组l
	for i := 1; i < length; i++ {
		l[i] = l[i-1] * nums[i-1]
	}

	// 填充数组r
	for i := length - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}

	for i := 0; i < length; i++ {
		result[i] = l[i] * r[i]
	}

	return result
}

// productExceptSelf2 除自身以外数组的乘积
// O(1) 的空间复杂度
func productExceptSelf2(nums []int) []int {

	length := len(nums)
	// 结果集
	result := make([]int, length)
	result[length-1] = 1

	// 想要O(1) 的空间复杂度 可以用result[i] 暂时保存i 右侧所有元素的乘积
	for i := length - 2; i >= 0; i-- {
		result[i] = result[i+1] * nums[i+1]
	}

	// 用l 表示i左侧所有元素的乘积，在遍历的过程中更新其值
	l := 1

	for i := 0; i < length; i++ {
		// 等号左侧的result[i] 题目要求的值，除i外其他元素的乘积
		// 等号右侧的result[i] 是指i右侧所尊元素的乘积
		// l 表示i左侧所有元素的乘积，初始为1，则result[0] = 1 * 0右侧所有元素的乘积
		result[i] = l * result[i]
		l = l * nums[i]
	}

	return result
}
