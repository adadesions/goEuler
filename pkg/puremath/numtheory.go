package puremath

import . "math"

func IsPrime(k int) bool {
	if k == 2 {
		return true
	} else if k < 2{
		return false
	}

	result := true
	lim := int(Ceil(Sqrt(float64(k))))
	for div := 2; div <= lim; div++ {
		if k % div == 0 {
			result = false
			break
		}
	}

	return result
}
