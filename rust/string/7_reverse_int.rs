

pub fn reverse(x: i32) -> i32 {
    let mut x = x;
    let mut ret = 0;

    while x != 0 {
        if ret > i32::MAX / 10 || ret < i32::MIN / 10{
            return 0
        }

        let carry = x % 10;
        x /= 10;
        ret = ret *  10 + carry;
    }
    ret 
}