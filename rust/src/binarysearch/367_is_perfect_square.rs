


// is_perfect_square 有效的完全平方数
// 0..num 二分查找
pub fn is_perfect_square(num: i32) -> bool {
    let num = num as i64;
    let (mut left, mut right) = (0_i64, num);

    while left <= right {
        let mid = left + ((right-left)>>1);

        let square = mid * mid;
        if square > num {
            right = mid -1;
        }else if square < num {
            left = mid + 1
        }else {
            return  true;
        }
    }
    false
}