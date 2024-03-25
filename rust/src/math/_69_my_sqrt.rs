/// x 的平方根
/// 二分法
pub fn my_sqrt(x: i32) -> i32 {
    let x = x as i64;

    let (mut l, mut r) = (0, x);

    while l <= r {
        let mid = l + ((r - l) >> 1);

        if mid * mid == x {
            return mid as i32;
        }

        if mid * mid < x {
            l = mid + 1;
        } else {
            r = mid - 1;
        }
    }

    // 如果不存在正好，l 在中间位置偏右
    if l * l > x {
        return l as i32 - 1;
    }
    l as i32
}
