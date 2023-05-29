

pub fn is_valid(s: String) -> bool {
    let s: Vec<char> = s.chars().collect();

    // 奇数
    if s.len() & 1 == 1 {
        return false
    }

    let mut stack = vec![];
    use std::collections::HashMap;

    let mut hmap = HashMap::new();
    hmap.insert('}', '{');
    hmap.insert(']', '[');
    hmap.insert(')', '(');

    for ch in s.iter() {
        // 右括号
        if hmap.contains_key(ch) {
            // 不匹配
            if stack.len() == 0 || hmap.get(ch)!= stack.get(stack.len()-1){
                return false
            }

            // 匹配 -> 出栈
            stack.pop();

        }else {
            stack.push(*ch)
        }
    }
    stack.len() == 0

}



pub fn is_valid2(s: String) -> bool {
    // match 匹配

    let mut stack = vec![];

    for ch in s.chars() {
        match ch {
            '[' | '{' | '(' => stack.push(ch),

            ']' => {if stack.pop() != Some('['){
                return false
            }},
            '}' => {if stack.pop() != Some('{') {
                return false
            }},
            ')' => {
                if stack.pop() != Some('(') {
                    return false
                }
            },
            _ => (),
        }
    }

    stack.is_empty()
}