package array

// canCompleteCircuit 加油站
// O(n) 时间复杂度
// 超时
func canCompleteCircuit(gas []int, cost []int) int {

	left := 0  // 剩余油量
	start := 0 // 起始出发点

	// 如果总油量 大于总的消耗量是有解的，否则无解
	totalGas, totalCost := 0, 0

	for i := 0; i < len(gas); i++ {

		totalCost += cost[i]
		totalGas += gas[i]

		left += gas[i] - cost[i] // 本次新增-本次要耗费的
		if left < 0 {
			start = i + 1
			left = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return start
}

// canCompleteCircuit 加油站
// O(n2) 时间复杂度
// 超时
func canCompleteCircuit2(gas []int, cost []int) int {

	n := len(gas)

	gas = append(gas, gas...) // 要走完一圈可以把gas 展开铺平
	cost = append(cost, cost...)

	for i := 0; i < n; i++ {

		isStart := true // 默认是可以循环的出发点

		left := 0 // 汽油的剩余量

		for j := i; j < i+n; j++ { //j 扫描n个点，能扫描完的话就是出发点

			left += gas[j] - cost[j]

			if left < 0 {
				isStart = false
				break
			}
		}

		if isStart {
			return i
		}

	}

	// 遍历完都没有符合条件的
	return -1
}
