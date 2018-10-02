package go_test

import "testing"

var z complex128
var cases = [][]complex128{
	{1 + 0i, 1 + 0i},
	{1 + 0i, 0 + 1i},
	{1 + 1i, 1 + 1i},
	{1 + 1i, 1 - 1i},
	{1 + 123456.987654i, 345678987654.5432 + 098754321.098765432i},
	{1 + 0.00000009876543i, 0.0000987654 - 0.000000000009876543456543i},
}

func approxEqual(a, b complex128) bool {
	dr := real(a) - real(b)
	di := imag(a) - imag(b)
	if dr < 0 {
		dr = -dr
	}
	if di < 0 {
		di = -di
	}
	if dr > 1e-6 || di > 1e-6 {
		return false
	}
	return true
}

func complexMultiplyBuiltin(a, b complex128) complex128 { return a * b }
func complexDivideBuiltin(a, b complex128) complex128   { return a / b }

func complexMultiplyExplicit(a, b complex128) complex128 {
	return complex(
		real(a)*real(b)-imag(a)*imag(b),
		real(a)*imag(b)+imag(a)*real(b),
	)
}

func complexDivideExplicit(a, b complex128) complex128 {
	// a/b = ab*/|b|**2, where b* is the complex conjugate of b and |b|**2 = b*cc(b)
	bmag := real(b)*real(b) + imag(b)*imag(b)
	return complex(
		(real(a)*real(b)+imag(a)*imag(b))/bmag,
		(-real(a)*imag(b)+imag(a)*real(b))/bmag,
	)
}

func TestComplexMultiply(t *testing.T) {
	for _, c := range cases {
		if res, out := complexMultiplyBuiltin(c[0], c[1]), complexMultiplyExplicit(c[0], c[1]); !approxEqual(out, res) {
			t.Errorf("%v*%v = %v, want %v", c[0], c[1], out, res)
		}
	}
}

func TestComplexDivide(t *testing.T) {
	for _, c := range cases {
		if res, out := complexDivideBuiltin(c[0], c[1]), complexDivideExplicit(c[0], c[1]); !approxEqual(out, res) {
			t.Errorf("%v/%v = %v, want %v", c[0], c[1], out, res)
		}
	}
}

func BenchmarkComplexMultiply(b *testing.B) {
	for _, c := range cases {
		b.Run("Direct", func(b *testing.B) {
			var zz complex128
			for i := 0; i < b.N; i++ {
				zz = complexMultiplyBuiltin(c[0], c[1])
			}
			z += zz
		})

		b.Run("Explicit", func(b *testing.B) {
			var zz complex128
			for i := 0; i < b.N; i++ {
				zz = complexMultiplyExplicit(c[0], c[1])
			}
			z += zz
		})
	}
}

func BenchmarkComplexDivide(b *testing.B) {
	for _, c := range cases {
		b.Run("Direct", func(b *testing.B) {
			var zz complex128
			for i := 0; i < b.N; i++ {
				zz = complexDivideBuiltin(c[0], c[1])
			}
			z += zz
		})

		b.Run("Explicit", func(b *testing.B) {
			var zz complex128
			for i := 0; i < b.N; i++ {
				zz = complexDivideExplicit(c[0], c[1])
			}
			z += zz
		})
	}
}
