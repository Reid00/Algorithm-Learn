/// 二进制求和
pub fn add_binary(a: String, b: String) -> String {
    let mut ans = String::new();

    let (a, b): (Vec<_>, Vec<_>) = (a.chars().collect(), b.chars().collect());

    let (len_a, len_b) = (a.len(), b.len());

    let n = len_a.max(len_b);

    let mut carry = 0;

    for i in 0..n {
        if i < len_a {
            let digit = a[len_a - i - 1].to_digit(10).unwrap();
            carry += digit;
        }

        if i < len_b {
            let digit = b[len_b - i - 1].to_digit(10).unwrap();
            carry += digit;
        }
        // 加法运算余数放到前面
        ans = format!("{}{}", carry % 2, ans);
        carry /= 2;
    }

    if carry > 0 {
        ans = format!("{}{}", 1, ans);
    }

    ans
}
