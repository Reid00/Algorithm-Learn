/// isSubsequence 判断子序列, s 是 t 的子序列
/// dp 为72.编辑距离做基础
pub fn is_subsequence(s: String, t: String) -> bool {
    // m <=n s 是否 是 t 的子序列
    // dp[i][j] 表示s以i-1 结尾和t 以j-1 结尾的相同最长子序列长度是dp[i][j]
    // if s[i-1] == t[j-1] => dp[i][j] = dp[i-1][j-1] + 1
    // if s[i-1] != t[j-1], 因为要判断s是否是t 的子序列，所以s 的下标不能减少
    // t 的可以缩短， 所以不相同时，t 可以剪短，删除一个字符，dp[i][j] = dp[i][j-1]

    let (m, n) = (s.len(), t.len());
    let (s, t) = (s.as_bytes(), t.as_bytes());

    let mut dp = vec![vec![0; n + 1]; m + 1];
    // 初始化, dp[0][0] 无意义，初始化为0
    // dp[i][0] 表示以下标i-1为结尾的字符串，与空字符串的相同子序列长度，所以为0. dp[0][j]同理。
    // 默认是0 无需初始化

    // 状态转移方程 遍历顺序
    for i in 1..m + 1 {
        for j in 1..n + 1 {
            if s[i - 1] == t[j - 1] {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            } else {
                // 字符串j 删除一个字符
                dp[i][j] = dp[i][j - 1];
            }
        }
    }
    dp[m][n] == m
}

/// isSubsequence 判断子序列, s 是 t 的子序列
/// 双指针
pub fn is_subsequence_(s: String, t: String) -> bool {
    // s 是否 是 t 的子序列
    let (s, t) = (s.as_bytes(), t.as_bytes());
    let (mut idx1, mut idx2) = (0, 0);
    let (m, n) = (s.len(), t.len());

    while idx1 < m && idx2 < n {
        if s[idx1] == s[idx2] {
            idx1 += 1;
        }

        idx2 += 1;
    }
    m == idx1
}
