package math

// maxPoints 直线上最多的点数
func maxPoints(points [][]int) int {
	res := 0

	for i := 0; i < len(points); i++ {
		// 建立一个key为两点的斜率，可能是浮点数
		hmap := make(map[float64]int)

		for j := 0; j < len(points); j++ {
			// 不能是同一个点组成的直线
			if i != j {
				hmap[lineSlop(points[i], points[j])]++
			}
		}

		for _, v := range hmap {
			if v > res {
				res = v
			}
		}
	}

	return res + 1 // hash表中没有统计点i 自身，所以需要加1
}

// lineSlop 斜率 float 计算除数为0 的时候，会转为inf
func lineSlop(a, b []int) float64 {
	return float64(a[1]-b[1]) / float64(a[0]-b[0]) // (ay-by)/(ax-bx)
}
