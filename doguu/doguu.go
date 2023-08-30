package doguu

func arithmeticMean(x []int) int {
	sum := 0

	for _, n := range x {
		sum += n
	}

	return sum / len(x)
}
