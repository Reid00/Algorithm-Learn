/// 基本计算器
pub fn calculate(s: String) -> i32 {
    // 栈
    // 计算值
    let mut res = 0;
    // 表示正负号  方便计算
    let mut sign = 1;
    let mut stack = vec![];

    let s = s.chars().collect::<Vec<char>>();

    let mut i = 0;
    while i < s.len() {
        let ch = s[i];

        match ch {
            '(' => {
                // 把做括号之前的计算值入栈，括号之前的符号也入栈
                stack.push(res);
                stack.push(sign);
                // 重置res 和sign 计算括号之内的值
                res = 0;
                sign = 1;
            }
            ')' => {
                // 把括号内的结果和之前的结果计算
                // (括号之前的符号
                let sign = stack.pop().unwrap();
                // (括号之前的计算值
                let prev_res = stack.pop().unwrap();
                // 计算值
                res = prev_res + sign * res;
            }
            '+' => sign = 1,
            '-' => sign = -1,
            '0' | '1' | '2' | '3' | '4' | '5' | '6' | '7' | '8' | '9' => {
                // 考虑后续是一个多位的整数，如 "123+1"
                let mut j = i;
                let mut num = 0;

                while j < s.len() && s[j].is_ascii_digit() {
                    num = num * 10 + (s[j] as i32 - '0' as i32);
                    j += 1;
                }

                res = res + sign * num;
                // j 在满足条件之后步进了一位, i 后续要+=1 会把非digit 的字符skip
                i = j - 1;
            }
            _ => (),
        }
        i += 1
    }

    res
}
