/// 编辑距离
pub fn min_distance(word1: String, word2: String) -> i32 {
    // dp[i][j] word1的前i个字符和word2的前j个字符的编辑距离。意思就是word1的前i个字符，变成word2的前j个字符，最少需要这么多步。
    // if word1[i-1] == word2[j-1]: dp[j][j] = dp[i-1][j-1]
    // if word1[i-] != word2[j-1]:
    // 有三种操作，1. 删: word1删除最后一个, dp[i][j] = dp[i-1][j] + 1
    // => word1删除一个元素，那么就是以下标i - 2为结尾的word1 与 j-1为结尾的word2的最近编辑距离 再加上一个操作
    // 删: word2 中删一个, dp[i][j] = dp[i][j-1] + 1
    // 2. 增： 增的本质和删一样，word1 删除一个，相当于word2 增加一个
    // 3. 改: word1替换word1[i - 1]，使其与word2[j - 1]相同，此时不用增删加元素, 只需要一次替换即可在i-2为下标的基础上转化
    // 即: dp[i][j] = dp[i-1][j-1] + 1
    // 总结, if word1[i-1]!=word2[j-1]=> dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1

    let (m, n) = (word1.len(), word2.len());
    let mut dp = vec![vec![0; n + 1]; m + 1];
    // 初始化dp
    dp[0][0] = 0;
    // 初始化第一列
    for i in 1..m + 1 {
        dp[i][0] = i;
    }
    // init first row
    for j in 1..n + 1 {
        dp[0][j] = j;
    }

    let (word1, word2): (Vec<_>, Vec<_>) = (word1.chars().collect(), word2.chars().collect());

    for i in 1..m + 1 {
        for j in 1..n + 1 {
            if word1[i - 1] == word2[j - 1] {
                dp[i][j] = dp[i - 1][j - 1];
            } else {
                dp[i][j] = dp[i - 1][j].min(dp[i][j - 1]).min(dp[i - 1][j - 1]) + 1;
            }
        }
    }
    dp[m][n] as i32
}
