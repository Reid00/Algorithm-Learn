/// 交错字符串
pub fn is_interleave(s1: String, s2: String, s3: String) -> bool {
    // dp[i][j]	表示s1 的前i个和s2 的前j个字符，能否构成s3的i+j 个字符
    // 前i个字符，对应的下标是i-1, 前j个字符是j-1, 第 i+j 个字符是 i + j -1
    // dp[i][j] = (dp[i-1][j] && s1[i-1]==s3[i+j-1]) or (dp[i][j-1] && s2[j-1] = s3[i+j-1])

    let (m, n) = (s1.len(), s2.len());
    if m + n != s3.len() {
        return false;
    }

    let mut dp = vec![vec![false; n + 1]; m + 1];
    // 初始化
    dp[0][0] = true;

    let (s1, s2, s3): (Vec<char>, Vec<char>, Vec<char>) = (
        s1.chars().collect(),
        s2.chars().collect(),
        s3.chars().collect(),
    );

    // 第一列初始化
    for i in 1..m + 1 {
        dp[i][0] = dp[i - 1][0] && s1[i - 1] == s3[i - 1];
    }

    // 第一行初始化
    for j in 1..n + 1 {
        dp[0][j] = dp[0][j - 1] && s2[j - 1] == s3[j - 1];
    }

    for i in 1..m + 1 {
        for j in 1..n + 1 {
            dp[i][j] = (dp[i - 1][j] && s1[i - 1] == s3[i + j - 1])
                || (dp[i][j - 1] && s2[j - 1] == s3[i + j - 1]);
        }
    }

    dp[m][n]
}
