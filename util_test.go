package jiwaifft

import "testing"

var floatSumple []float64 = []float64{1, 1.5, 0, -1.5, 4}
var complexSumple []complex128 = []complex128{1 + 0i, 1.5 + 0i, 0i, -1.5 + 0i, 4 + 0i}

func TestF2C1d(t *testing.T) {
	res := F2C1d(floatSumple)
	lx := len(floatSumple)
	for i := 0; i < lx; i++ {
		if res[i] == complexSumple[i] {
			continue
		} else {
			t.Fatalf("i=%v in=%v,out=%v,correct=%v", i, floatSumple[i], res[i], complexSumple[i])
		}
	}
}
