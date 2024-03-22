/// 颠倒二进制位
/// 位运算
pub fn reverse_bits(x: u32) -> u32 {
    let mut res = 0_u32;
    let mut x = x;

    for _ in 0..32 {
        res = (res << 1) | (x & 1);
        x >>= 1;
    }
    res
}
