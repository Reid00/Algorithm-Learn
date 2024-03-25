/// 阶乘后的0
pub fn trailing_zeroes(n: i32) -> i32 {
    let mut cnt = 0;
    let mut n = n;
    while n > 0 {
        n /= 5;
        cnt += n;
    }
    cnt
}
