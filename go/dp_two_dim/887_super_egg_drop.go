package dptwodim

// superEggDrop 鸡蛋掉落
// 请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？
// k 个鸡蛋，n层建筑
func superEggDrop(k int, n int) int {
	// dp[i][j] 表示剩余i个鸡蛋，还有j次仍鸡蛋的机会时，能确定的最高楼层是 dp[i][j] （即最多可以测出多少层）
	// 比如，1个鸡蛋，7次机会，最高可以确定的楼层是7(从一楼开始一层层的仍)
	// 如果鸡蛋没有碎，此时需要往上走, 还有j-1次机会，之后可以最多搜索的楼层数是dp[i][j-1]
	// 如果鸡蛋碎了，此时需要往下走, 剩余i-1个鸡蛋，和j-1次机会，后面最多可以搜索的最多楼层数是dp[i-1][j-1]
	// 加上本层的x，则状态转移方程为: dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
	// 在本题中，要求我们最坏情况下: 当所有鸡蛋都用了，一定要 dp[k][j] >= n 即一定要测出n层

	dp := make([][]int, k+1)
	// 因为要求仍的次数j 并不知道，实时其最多为n次 (一个鸡蛋从一楼一层一层的仍)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化
	// 鸡蛋为0， 次数为0 能确定的楼层数都是0
	// 第一列和第一行都是0

	// 确定遍历顺序
	// 先遍历仍的次数，保证每个鸡蛋都可以扔一次
	// 如果反过来，先遍历鸡蛋，则会出现一个鸡蛋仍n次，实际上这个是不可能的
	for j := 1; j < n+1; j++ {

		for i := 1; i < k+1; i++ {

			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1

			if dp[i][j] >= n {
				return j
			}
		}
	}

	return n
}

func superEggDrop2(k int, n int) int {

	if k == 0 || n == 0 {
		return 0
	}

	// 只有一个鸡蛋一层层的向上试，最坏情况是f=n
	if k == 1 {
		return n
	}
	// 只有一层楼，开始扔在1楼即可知道结果，碎了f=0,不碎f=1
	if n == 1 {
		return 1
	}
	// 最多仍的次数是n，定义一个最大值，后续更新结果
	minDrop := n + 1

	for i := 1; i < n+1; i++ {
		// 递归
		minDrop = min(minDrop, max(superEggDrop2(k, n-i)), superEggDrop2(k-1, i-1)) + 1
	}

	return minDrop
}
