/// 回文子串     数量
/// 中心扩展算法 O(n2) 的时间复杂度 + O(1) 的空间
pub fn count_substrings(s: String) -> i32 {
    let mut res = 0;
    let s: Vec<char> = s.chars().collect();

    let n = s.len() as i32;

    for i in 0..n {
        res += expand_from_center(&s, i, i); // 以i为中心向两边扩散
        res += expand_from_center(&s, i, i + 1); // 以i,i+1 为中心向两边扩散
    }

    res as i32
}

pub fn expand_from_center(s: &[char], mut i: i32, mut j: i32) -> usize {
    let mut res = 0;
    let n = s.len() as i32;

    while i >= 0 && j < n && s[i as usize] == s[j as usize] {
        i -= 1;
        j += 1;
        res += 1;
    }
    res
}

/// 回文子串     数量
/// dp O(n2) 的时间复杂度 + O(n2) 的空间复杂度
pub fn count_substrings2(s: String) -> i32 {
    let mut res = 0;
    let n = s.len();
    let mut dp = vec![vec![false; n]; n];

    let s: Vec<char> = s.chars().collect();

    for i in (0..n).rev() {
        // 由dp 的定义可知，j >=i 所以只需要填充矩阵的右上方
        for j in i..n {
            if s[i] == s[j] {
                if j - i < 2 {
                    dp[i][j] = true;
                    res += 1;
                } else if dp[i + 1][j - 1] {
                    dp[i][j] = dp[i + 1][j - 1];
                    res += 1;
                }
            }
        }
    }

    res
}
