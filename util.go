package jiwaifft

import "math/cmplx"

// F2C1d is convert []float64 to []complex128
func F2C1d(x []float64) []complex128 {
	lx := len(x)
	res := make([]complex128, lx, lx)
	for i := 0; i < lx; i++ {
		res[i] = complex(x[i], 0)
	}
	return res
}

// C2Power1d is convert []Abs(complex128) to []float64
func C2Power1d(x []complex128) []float64 {
	lx := len(x)
	res := make([]float64, lx, lx)
	for i := 0; i < lx; i++ {
		res[i] = cmplx.Abs(x[i])
	}
	return res
}

// C2Real1d is convert []Real(complex128) to []float64
func C2Real1d(x []complex128) []float64 {
	lx := len(x)
	res := make([]float64, lx, lx)
	for i, val := range x {
		res[i] = real(val)
	}
	return res

}
