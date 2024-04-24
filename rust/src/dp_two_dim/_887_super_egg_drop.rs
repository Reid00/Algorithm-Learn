/// 鸡蛋掉落
pub fn super_egg_drop(k: i32, n: i32) -> i32 {
    // dp[i][j] 表示剩余i个鸡蛋，还有j次仍鸡蛋的机会时，能确定的最高楼层是 dp[i][j] （即最多可以测出多少层）
    // 比如，1个鸡蛋，7次机会，最高可以确定的楼层是7(从一楼开始一层层的仍)
    // 2个鸡蛋，2次机会，最高可以确定的楼层是3
    // 假设现在扔在x层，
    // 如果鸡蛋没有碎，此时需要往上走, 还有j-1次机会，之后可以最多搜索的楼层数是dp[i][j-1]
    // 如果鸡蛋碎了，此时需要往下走, 剩余i-1个鸡蛋，和j-1次机会，后面最多可以搜索的最多楼层数是dp[i-1][j-1]
    // 加上本层的x，则状态转移方程为: dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1


    let (k, n) = (k as usize, n as usize);

    let mut dp = vec![vec![0; n + 1]; k + 1];

    // 初始化
    // 0 个鸡蛋或者0次机会，最高能确定的楼层数都是0

    // 先遍历仍的次数，保证每个鸡蛋都可以扔一次
    // 如果反过来，先遍历鸡蛋，则会出现一个鸡蛋仍n次，实际上这个是不可能的
    for j in 1..n + 1 {
        for i in 1..k + 1 {
            dp[i][j] = dp[i][j - 1] + dp[i - 1][j - 1] + 1;

            if dp[i][j] >= n {
                return j as i32;
            }
        }
    }

    n as i32
}
