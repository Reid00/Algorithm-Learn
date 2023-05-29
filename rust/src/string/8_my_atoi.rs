


pub fn my_atoi(s: String) -> i32 {
    let mut s: Vec<char> = s.trim().chars().collect();
    if s.len() == 0 {
        return 0
    }

    let mut result: i64 = 0;
    let mut negative = false;

    if s[0] == '-' {
        negative = true;
        s.remove(0);
    }else if s[0] == '+' {
        s.remove(0);
    }

    for &ch in s.iter() {
        if ch >= '0' && ch <= '9' {
            let ch = ch.to_digit(10).unwrap() as i64;
            result = result * 10 + ch;

            if result >= 1 << 31 {
                result = 1 << 31;
                break;
            }
        }else {
            break;
        }
    }

    println!("result {}", result);

    // 处理边界问题
    if negative {
        result = - result;
    }else {
        if result == 1 << 31 {
            result = (1<<31) - 1;
            println!("max: {}", result);
        }
    }
    result as i32
}
