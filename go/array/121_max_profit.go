package array

import "math"

func maxProfit(prices []int) int {

	// 假设自己在最低点买入，遍历所有价格，计算哪天的profit 最高

	minVal := math.MaxInt64 // 最低点的股票

	maxProfit := 0 // 最大的收益值

	for _, v := range prices {
		if v < minVal {
			minVal = v
		} else if v-minVal > maxProfit {
			maxProfit = v - minVal
		}
	}

	return maxProfit

}
