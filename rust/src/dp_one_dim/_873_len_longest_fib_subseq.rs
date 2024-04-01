/// 最长的斐波那契子序列的长度
pub fn len_longest_fib_subseq(arr: Vec<i32>) -> i32 {
    let n = arr.len();
    let mut ans = 0;

    let mut dp = vec![vec![2; n]; n];

    use std::collections::HashMap;
    let mut hmap = HashMap::new();

    for (i, &v) in arr.iter().enumerate() {
        hmap.insert(v, i);
    }

    for i in 0..n {
        for j in i + 1..n {
            let v = arr[i] + arr[j];
            if hmap.contains_key(&v) {
                let &idx = hmap.get(&v).unwrap();
                dp[j][idx] = dp[j][idx].max(dp[i][j] + 1);
                ans = ans.max(dp[j][idx]);
            }
        }
    }

    if ans >= 3 {
        return ans;
    }
    return 0;
}
