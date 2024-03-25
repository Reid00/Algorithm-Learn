package math

// myPow Pow(x, n)
func myPow(x float64, n int) float64 {
	// 快速幂法
	if n > 0 {
		return quickMu(x, n)
	}

	return 1 / quickMu(x, -n)
}

func quickMu(x float64, n int) float64 {

	if n == 0 {
		return 1
	}

	y := quickMu(x, n/2)

	// 偶数
	if n%2 == 0 {
		return y * y
	}
	// 奇数
	return y * y * x
}
