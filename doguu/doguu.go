package doguu

func arithmeticMean(x []int) float64 {
	sum := 0

	for _, n := range x {
		sum += n
	}

	return float64(sum) / float64(len(x))
}

func geometricMean(x []float64) float64 {

	return float64(1)
}
