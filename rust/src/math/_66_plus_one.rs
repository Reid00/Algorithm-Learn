/// 加一
pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
    let mut digits = digits;

    let n = digits.len();

    for i in (0..n).rev() {
        // 从后面的末尾位置 + 1
        digits[i] += 1;

        digits[i] = digits[i] % 10;

        // 末位+1 后 如果不为10，则直接返回，否则前一位+1
        if digits[i] != 0 {
            return digits;
        }
    }

    // 如果全为9, 首位加一位
    let mut digits = vec![0; n + 1];
    digits[0] = 1;

    digits
}
