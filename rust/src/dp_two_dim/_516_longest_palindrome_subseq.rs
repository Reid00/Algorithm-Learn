/// 最长回文子序列 的长度 （注意是子序列不是子串）
pub fn longest_palindrome_subseq(s: String) -> i32 {
    // dp[i][j] 表示s[i,j] 最长回文子序列的最大长度是dp[i][j] (包含i,j)
    // if s[i]==s[j], dp[i][j] = dp[i+1][j-1] + 2
    // if s[i] != s[j], 说明s[i]和s[j]的同时加入 并不能增加[i,j]区间回文子序列的长度，那么分别加入s[i]、s[j]看看哪一个可以组成最长的回文子序列。
    // 1. 加入s[i], dp[i][j] = dp[i][j-1]
    // 2. 加入s[j], dp[i][j] = dp[+1][j]
    // so. dp[i][j] = max(dp[i][j-1], dp[i+1][j])
    // thus, dp[i][j] = dp[i+1][j-1] +2 Or max(dp[i][j-1], dp[i+1][j])

    let n = s.len();
    let s: Vec<char> = s.chars().collect();

    let mut dp = vec![vec![0; n]; n];

    for i in 0..n {
        dp[i][i] = 1;
    }

    for i in (0..n).rev() {
        for j in i + 1..n {
            if s[i] == s[j] {
                dp[i][j] = dp[i + 1][j - 1] + 2;
            } else {
                dp[i][j] = dp[i][j - 1].max(dp[i + 1][j]);
            }
        }
    }

    dp[0][n - 1]
}
