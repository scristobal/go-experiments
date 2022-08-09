package sqrt

import (
	"math/rand"
	"testing"
)

func Benchmark(b *testing.B) {

	var N = 1000 * rand.Intn(10_000)

	b.Run("bounded", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Bounded(N)
		}
	})

	b.Run("newton", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Newton(N)
		}
	})

	b.Run("divide", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Divide(N)
		}
	})
}
