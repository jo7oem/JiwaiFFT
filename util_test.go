package jiwaifft

import (
	"math"
	"testing"
)

var floatSumple = []float64{1, 1.5, 0, -1.5, 4, 1, 1.5, 0, -1.5, 4, 1, 1.5, 0, -1.5, 4, 1, 1.5, 0, -1.5, 4}
var complexSumple = []complex128{1 + 0i, 1.5 + 0i, 0i, -1.5 + 0i, 4 + 0i, 1 + 0i, 1.5 + 0i, 0i, -1.5 + 0i, 4 + 0i, 1 + 0i, 1.5 + 0i, 0i, -1.5 + 0i, 4 + 0i, 1 + 0i, 1.5 + 0i, 0i, -1.5 + 0i, 4 + 0i}

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

func TestC2Power1d(t *testing.T) {
	res := C2Power1d(complexSumple)
	lx := len(complexSumple)
	for i := 0; i < lx; i++ {
		if res[i] == math.Abs(floatSumple[i]) {
			continue
		} else {
			t.Fatalf("i=%v in=%v,out=%v,correct=%v", i, complexSumple[i], res[i], math.Abs(floatSumple[i]))
		}
	}
}

func TestC2Real1d(t *testing.T) {
	res := C2Real1d(complexSumple)
	lx := len(complexSumple)
	for i := 0; i < lx; i++ {
		if res[i] == floatSumple[i] {
			continue
		} else {
			t.Fatalf("i=%v in=%v,out=%v,correct=%v", i, complexSumple[i], res[i], floatSumple[i])
		}
	}
}
func TestC2Imag1d(t *testing.T) {
	res := C2Imag1d(complexSumple)
	lx := len(complexSumple)
	for i := 0; i < lx; i++ {
		if res[i] == 0 {
			continue
		} else {
			t.Fatalf("i=%v in=%v,out=%v,correct=%v", i, complexSumple[i], res[i], 0)
		}
	}
}
