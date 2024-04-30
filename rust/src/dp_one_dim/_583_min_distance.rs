/// 两个字符串的删除操作
/// 给定两个单词 word1 和 word2 ，返回使得 word1 和  word2 相同所需的最小步数。
pub fn min_distance(word1: String, word2: String) -> i32 {
    // dp[i][j] 表示word1 以i-1 为下标的字符串，转化为以j-1为下标的字符串word2,最小删除的步数是dp[i][j]
    // if word1[i-1] == word2[j-1] => dp[i][j] = dp[i-1][j-1]
    // if word1[i-1] != word2[j-1] => dp[i][j] = min(dp[i-1][j] + 1, dp[i][j-1] + 1)
    // word1 前i-1个下标中删除一个字符，转化为word2[:j] 或者word1不懂，word2前j-1 个字符删除一个字符转为word1[:i]

    let (m, n) = (word1.len(), word2.len());
    let (word1, word2) = (word1.as_bytes(), word2.as_bytes());

    let mut dp = vec![vec![0; n + 1]; m + 1];

    // 初始化
    for i in 0..m + 1 {
        // word1 下标为i-1的 i个字符 转化为空字符串，需要删除i次
        dp[i][0] = i;
    }

    for j in 0..n + 1 {
        dp[0][j] = j;
    }

    for i in 1..m + 1 {
        for j in 1..n + 1 {
            if word1[i - 1] == word2[j - 1] {
                dp[i][j] = dp[i - 1][j - 1];
            } else {
                dp[i][j] = dp[i - 1][j].min(dp[i][j - 1]) + 1;
            }
        }
    }

    dp[m][n] as i32
}
