package array

func maxProfit2(prices []int) int {

	// 贪心算法，由局部最优 推出全局最优

	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit

}
