/// 不同的子序列
/// s 序列中 t 出现的个数
pub fn num_distinct(s: String, t: String) -> i32 {
    // dp[i][j] 表示以i-1 下标为结尾的s序列中，出现以j-1 下标为结尾的子序列t 的个数是dp[i][j]
    // if s[i-1] == t[j-1]，有两种可能，不管最后一个字符 dp[i][j] = dp[i-1][j-1]
    // 第二个情况是，在s 中不用以i-1 下标为结尾，删除一个元素继续匹配，可以用i-2 即dp[i][j] = dp[i-1][j]
    // s 中前i-2 个字符中，有多少个t 中前j-1 个字符, 比如s=bagg t=bag dp[4][3] 可以用s[0][1][2] 匹配t[3]
    // so dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
    // if s[i-1] != t[j-1], dp[i][j] = dp[i-1][j]

    let (m, n) = (s.len(), t.len());
    let (s, t) = (s.as_bytes(), t.as_bytes());

    let mut dp = vec![vec![0; n + 1]; m + 1];

    // 初始化
    // 初始化行 s 中前i-1个字符，转化为空字符串的个数 为1 （删除所有字符）
    for i in 0..m + 1 {
        dp[i][0] = 1;
    }

    for j in 0..n + 1 {
        dp[0][j] = 0;
    }

    dp[0][0] = 1;

    for i in 1..m + 1 {
        for j in 1..n + 1 {
            if s[i - 1] == t[j - 1] {
                dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j];
            } else {
                dp[i][j] = dp[i - 1][j];
            }
        }
    }

    dp[m][n]
}
