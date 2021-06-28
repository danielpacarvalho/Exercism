// Package triangle check what kind of triangle it is
package triangle

import "math"

// Notice KindFromSides() returns this type.
type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides calculates what kind of triangle it is and return the iota equivalent for its kind
func KindFromSides(a, b, c float64) Kind {
	if isTriangle(a, b, c) {
		if a == b && b == c {
			return Equ
		}
		if a != b && b != c && a != c {
			return Sca
		}
		if a == b || b == c || a == c {
			return Iso
		}
	}
	return NaT
}

func isTriangle(a, b, c float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}
	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return false
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if a+b < c {
		return false
	}
	if b+c < a {
		return false
	}
	if a+c < b {
		return false
	}
	return true
}
