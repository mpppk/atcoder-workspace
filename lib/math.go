package lib

import "math"

// PowAAA は、xのy乗を返します.
func PowAAA(x, y AAA) AAA {
	return AAA(math.Pow(float64(x), float64(y)))
}
