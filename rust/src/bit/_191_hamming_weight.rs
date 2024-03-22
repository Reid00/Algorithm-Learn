/// 位 1 的个数
pub fn hammingWeight(n: u32) -> i32 {
    let mut res = 0;
    let mut n = n;

    while n > 0 {
        n &= n - 1;
        res += 1;
    }
    res
}
