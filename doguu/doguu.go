package doguu

import "math"

func arithmeticMean(x []int) float64 {
	sum := 0

	for _, v := range x {
		sum += v
	}

	return float64(sum) / float64(len(x))
}

func geometricMean(x []float64) float64 {
	n := 1.0

	for _, v := range x {
		n *= v
	}

	return math.Pow(n, 1.0/float64(len(x)))
}
