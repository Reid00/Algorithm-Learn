/// 中心扩散算法
/// Time: O(n2), Space: O(n2)
pub fn longest_palindrome(s: String) -> String {
    let n = s.len() as i32;
    if n < 2 {
        return s;
    }

    let s: Vec<char> = s.chars().collect();
    let (mut l, mut r) = (0, 0);

    for i in 0..n {
        let (s1, e1) = expand(&s, i, i); // 以i 为中心进行扩散
        let (s2, e2) = expand(&s, i, i + 1); // 以i,i+1 为中心进行扩散

        if e1 - s1 > r - l {
            (l, r) = (s1, e1);
        }

        if e2 - s2 > r - l {
            (l, r) = (s2, e2);
        }
    }

    s[l as usize..r as usize + 1].iter().collect()
}

pub fn expand(s: &[char], mut i: i32, mut j: i32) -> (i32, i32) {
    while i >= 0 && j < s.len() as i32 && s[i as usize] == s[j as usize] {
        i -= 1;
        j += 1;
    }

    (i + 1, j - 1)
}

/// 最长回文子串
/// Time: O(n2), Space: O(n2)
pub fn longest_palindrome2(s: String) -> String {
    let n = s.len();
    if n < 2 {
        return s;
    }

    let mut dp = vec![vec![false; n]; n];

    let s: Vec<char> = s.chars().collect();
    let (mut max_start, mut max_len) = (0, 0);

    // dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
    // 由于(i,j) 依赖 (i+1,j-1) 在左下角，所以遍历顺序，从下到上，从左到右
    for i in (0..n).rev() {
        for j in i..n {
            if s[i] == s[j] {
                if j - i < 2 {
                    dp[i][j] = true;
                } else {
                    dp[i][j] = dp[i + 1][j - 1];
                }
            }

            if dp[i][j] && j - i + 1 > max_len {
                max_start = i;
                max_len = j - i + 1;
            }
        }
    }

    s[max_start..max_start + max_len].iter().collect()
}
