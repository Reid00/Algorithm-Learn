/// 基本计算器 II
pub fn calculate(s: String) -> i32 {
    // 类似224 有 + - * / 没有 括号
    // 224 只有 + - 可以没有碰到括号的情况下，直接计算结果
    // 本体 没有括号，有 * /  所以要先把所有结果放到栈中，碰到 * / 先计算 ，最后遍历stack 用 + 计算结果

    let mut res = 0;
    // 因为乘除需要知道left right 才能计算，所以做判断时，把当前数字做right 根据前面的字符是否是
    // * / 来做计算
    let mut pre_sign = '+';

    let mut stack = vec![];

    // 遍历时 要压栈的数字
    let mut num = 0;

    for (i, ch) in s.char_indices() {
        if ch.is_ascii_digit() {
            num = num * 10 + ch as i32 - '0' as i32;
        }

        if !ch.is_ascii_digit() && ch != ' ' || i == s.len() - 1 {
            match pre_sign {
                '+' => stack.push(num),
                '-' => stack.push(-num),
                '*' => {
                    let prev_num = stack.pop().unwrap();
                    num = prev_num * num;
                    stack.push(num);
                }
                '/' => {
                    stack.last_mut().map(|prev| *prev /= num);
                    // stack
                    //     .last_mut()
                    //     .map(|prev| *prev = prev.checked_div(num).unwrap_or(0));
                }
                _ => (),
            }
            // 如果当前字符不是数字，更新pre_sign
            pre_sign = ch;
            num = 0
        }
    }

    for i in stack {
        res += i;
    }

    res
}
