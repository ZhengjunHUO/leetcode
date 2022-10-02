package extra0003

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	lenth := 5
	value := []int{10,2,8,6,4}
	cost := []int{5,5,3,3}
	risk := []float32{0,1,0.5,0.15}
	expected := []float32{25,9,17,20.10825}

	for i := range cost {
		rslt := MaxProfit(value, lenth, cost[i], risk[i])
		if rslt != expected[i] {
			t.Errorf("Wanted %v, but get %v", expected[i], rslt)
		}
	}
}
