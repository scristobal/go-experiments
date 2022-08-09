package sqrt

import (
	"math"
	"math/rand"
	"testing"
)

func TestNewtonInt(t *testing.T) {
	n := NewtonInt(24)
	if n != 4 {
		t.Errorf("NewtonInt(24) was %d should be 4 ", n)
	}
}

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

	b.Run("newton-int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			NewtonInt(N)
		}
	})

	b.Run("divide", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Divide(N)
		}
	})

	b.Run("built-in", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			math.Sqrt(float64(N))
		}
	})
}
