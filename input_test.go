package jiwaifft

import (
	"fmt"
	"testing"
)

func TestReadCSVData1d(t *testing.T) {
	res, err := ReadCSVData1d(`testData/PictDatahead10.csv`, 1)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	if lx := len(res); lx != 1 {
		t.Fatalf("[][]float64 fail.[%v][]float64", lx)
	}
	for i, val := range res[0] {
		fmt.Printf("[%v],%v\n", i, val)
	}
}
