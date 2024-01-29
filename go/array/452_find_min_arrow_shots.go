package array

import "slices"

// findMinArrowShots 用最少数量的箭引爆气球
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	// 先排序 根据x[1] 升序排序, 因为元素1决定了箭的最大右侧位置
	slices.SortStableFunc(points, func(a, b []int) int {
		return a[1] - b[1]
	})

	// 确认需要几只箭，默认把箭放到points[0][1] 的位置，观察这个位置包含了几个区域
	maxRight := points[0][1]
	res := 1

	for _, v := range points {
		// 目前这只箭不能引爆v 这个气球，增加一支箭，让其等于下一个points 元素的最右侧
		if v[0] > maxRight {
			maxRight = v[1]
			res++
		}
	}

	return res
}
