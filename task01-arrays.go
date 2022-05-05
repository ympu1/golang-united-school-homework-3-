package homework

func average(input []float32) (result float32) {
	//Place your code here
	var totalSum float32
	for _, val := range input {
		totalSum += val
	}

	return totalSum / float32(len(input))
}
