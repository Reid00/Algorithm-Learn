/// 逆波兰表达式求值
pub fn eval_rpn(tokens: Vec<String>) -> i32 {
    // 栈模拟

    let mut stack = vec![];

    for v in tokens {
        match v.parse::<i64>() {
            Ok(num) => stack.push(num),

            Err(_) => {
                let right = stack.pop().unwrap();
                let left = stack.pop().unwrap();

                let mut res = 0_i64;

                match v.as_str() {
                    "+" => res = left + right,
                    "-" => res = left - right,
                    "*" => res = left * right,
                    "/" => res = left / right,
                    _ => (),
                }
                stack.push(res)
            }
        }
    }

    stack[0] as i32
}
