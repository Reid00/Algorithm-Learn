

pub fn is_subsequence(s: String, t: String) -> bool {
    // s 是否 是 t 的子序列
    let (s, t) = (s.as_bytes(), t.as_bytes());
    let (mut idx1, mut idx2) = (0, 0);
    let (m, n) = (s.len(), t.len());

    while idx1 < m && idx2 < n {
        if s[idx1] == t[idx2] {
            idx1 += 1;
        }
        idx2 += 1;
    }
    idx1 == m
}
