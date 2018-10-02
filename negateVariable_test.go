package go_test

import "testing"

func BenchmarkNegateVariable(b *testing.B) {
	x := complex(1232, -432.3)
	b.Run("x=-x", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x = -x
		}
	})
	b.Run("x*=-1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x *= -1
		}
	})
}
