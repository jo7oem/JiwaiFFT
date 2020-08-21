package jiwaifft

import "math/cmplx"

func F2C1d(x []float64) []complex128 {
	lx := len(x)
	res := make([]complex128, lx, lx)
	for i := 0; i < lx; i++ {
		res[i] = complex(x[i], 0)
	}
	return res
}

func C2F1d(x []complex128) []float64 {
	lx := len(x)
	res := make([]float64, lx, lx)
	for i := 0; i < lx; i++ {
		res[i] = cmplx.Abs(x[i])
	}
	return res
}
