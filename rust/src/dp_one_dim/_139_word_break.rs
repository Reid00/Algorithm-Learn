/// 单词拆分
pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
    let n = s.len();
    let mut dp = vec![false; n + 1];
    dp[0] = true;

    for i in 0..n {
        for j in i + 1..n + 1 {
            if dp[i] && word_dict.iter().any(|w| w == &s[i..j]) {
                dp[j] = true;
            }
        }
    }
    dp[n]
}
