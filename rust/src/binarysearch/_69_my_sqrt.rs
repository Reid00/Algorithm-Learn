
pub fn my_sqrt(x: i32) -> i32 {
    let x = x as i64;
    let (mut left, mut right) = (0, x);

    while left <= right {
        let mid = left + ((right-left)>>1);

        let square = mid * mid ;

        if square == x {
            return mid as i32
        }

        if square > x {
            right = mid -1;
        }else {
            left = mid + 1;
        }
    }

    if left * left > x {
        return (left -1) as i32
    }

    left as i32
}