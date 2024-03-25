/// 回文数 121 是， 123 不是
pub fn is_palindrome(x: i32) -> bool {
    if x < 0 {
        return false;
    }

    let mut origin = x;
    let mut res = 0;

    while origin > 0 {
        let carry = origin % 10;
        res = res * 10 + carry;

        origin /= 10;
    }

    x == res
}
