package extra

import (
	"testing"
)

func TestXorProduct(t *testing.T) {
	a, b, want := []int{5, 12}, []int{8, 21}, []int{12, 1}
	for i := range a {
		if temp := XorProductBF(a[i], b[i]); temp != want[i] {
			t.Errorf("XorProduct(%d, %d) return %d but want %d", a[i], b[i], temp, want[i])
		}
	}
}
