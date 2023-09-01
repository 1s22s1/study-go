package doguu

import "math"

func arithmeticMean(x []float64) float64 {
	n := 0.0

	for _, v := range x {
		n += v
	}

	return float64(n) / float64(len(x))
}

func geometricMean(x []float64) float64 {
	n := 1.0

	for _, v := range x {
		n *= v
	}

	return math.Pow(n, 1.0/float64(len(x)))
}

func harmonicMean(x []float64) float64 {
	n := 0.0

	for _, v := range x {
		n += (1.0 / v)
	}

	return float64(len(x)) / n
}

func variance(x []float64) float64 {
	n := 0.0
	mean := arithmeticMean(x)

	for _, v := range x {
		n += (v - mean) * (v - mean)
	}

	return n / float64(len(x))

}
