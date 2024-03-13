/// 字母大小写全排列
pub fn letter_case_permutation(s: String) -> Vec<String> {
    if s.is_empty() {
        return vec![];
    }

    let mut res = vec![];

    let mut path = "".to_string();
    let s: Vec<_> = s.chars().collect();

    backtrack(&s, 0, &mut res, &mut path);
    res
}

fn backtrack(s: &[char], idx: usize, res: &mut Vec<String>, path: &mut String) {
    if idx == s.len() {
        res.push(path.clone());
        return;
    }

    path.push(s[idx]);
    backtrack(s, idx + 1, res, path);
    path.pop();

    if s[idx].is_ascii_lowercase() {
        path.push(s[idx].to_ascii_uppercase());
        backtrack(s, idx + 1, res, path);
        path.pop();
    } else if s[idx].is_ascii_uppercase() {
        path.push(s[idx].to_ascii_lowercase());
        backtrack(s, idx + 1, res, path);
        path.pop();
    }
}
