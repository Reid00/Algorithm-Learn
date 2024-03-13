use std::collections::HashSet;

/// 拆分字符串使唯一子字符串的数目最大
pub fn max_unique_split(s: String) -> i32 {
    if s.is_empty() {
        return 0;
    }

    // 子串最终的数目
    let mut res = 0;

    // 搜索过程中的字符串子集，不能重复要用集合
    let mut path = HashSet::new();
    let s: Vec<char> = s.chars().collect();
    backtrack(&s, 0, 0, &mut path, &mut res);
    res
}

fn backtrack(s: &[char], idx: usize, count: i32, path: &mut HashSet<Vec<char>>, res: &mut i32) {
    // 剪枝
    if s.len() - idx + count as usize <= *res as usize {
        return;
    }

    if idx >= s.len() {
        *res = (*res).max(count);
        return;
    }

    for i in idx..s.len() {
        let sub_str = &s[idx..i + 1].to_vec();
        if !path.contains(sub_str) {
            path.insert(sub_str.clone());
            backtrack(s, i + 1, count + 1, path, res);
            path.remove(sub_str);
        }
    }
}
