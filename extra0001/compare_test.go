package extra

import (
	"testing"
)

func BenchmarkXorproduct(b *testing.B) {
	for i:=0; i<b.N; i++ {
		XorProduct(1000000, 1100000)
	}
}

func BenchmarkXorProductBF(b *testing.B) {
	for i:=0; i<b.N; i++ {
		XorProductBF(1000000, 1100000)
	}
}
