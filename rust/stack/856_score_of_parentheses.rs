

// score_of_parentheses 栈的用法，碰到右括号，计算这对括号对
// 上对括号的贡献值
pub fn score_of_parentheses(s: String) -> i32 {

    let mut ans = vec![0];

    for ch in s.chars() {
        if ch == '(' {
            ans.push(0);
        }else {
            // 先删除对应的左括号, 然后计算内部的分数
            
            if let Some(v)= ans.pop() {
                // 更新 这对删除的 平衡左括对ans 的分数影响 +=
                if let Some(last) = ans.last_mut() {
                    // *last += std::cmp::max(2 * v, 1);
                    *last += 1.max(2 * v)
                }
            }
        }
    }

    ans[0]
}

// 当作树 => 根节点的值是 每个叶子节点值 与其深度的幂的总和。
pub fn score_of_parentheses2(s: String) -> i32 {

    let mut ans = 0;
    let mut depth = 0;

    let chars: Vec<char> = s.chars().collect();

    for (i, &ch) in chars.iter().enumerate(){
        if ch == '(' {
            depth += 1;
        }else {
            depth -= 1;
            
            // 这对平衡括号的值，等于2 的其深度的幂次方
            if ch == ')' && chars[i-1] == '(' {
                ans += 1 << depth;
            }
        }
    }

    ans
}