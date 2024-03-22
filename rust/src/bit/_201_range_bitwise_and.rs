/// 数字范围按位与
pub fn range_bitwise_and(left: i32, right: i32) -> i32 {
    // METHOD 1: 遍历从 left - right 都包含，然后依次相与，但是如果特别大，就会超时
    // METHOD 2: 因为left 和 right 之间的元素是依次递增的，二进制表示就是每位相差1
    // 因此若干相与的时候，后缀部分由于相差1，有0的最终都会变成0
    // 最终变成了返回left, right 的二进制前缀部分
    // x & (x-1) 可以消除右边的1，因此 大数right 逐渐去掉右边的1 等于left 时即是结果

    let mut right = right;

    while left < right {
        right = right & (right - 1)
    }

    right
}
